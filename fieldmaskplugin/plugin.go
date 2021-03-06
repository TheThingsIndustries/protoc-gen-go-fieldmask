// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package fieldmaskplugin

import (
	"fmt"
	"strings"
)

// FieldSet is a set of fields.
type FieldSet map[string]struct{}

// Add adds the fields to the set. Returns true iff all fields were added (not already in the set).
func (fs FieldSet) Add(fields ...string) bool {
	ok := true
	for _, field := range fields {
		if !fs.add(field) {
			ok = false
		}
	}
	return ok
}

func (fs FieldSet) add(field string) bool {
	if _, found := fs[field]; found {
		return false
	}
	fs[field] = struct{}{}
	return true
}

// Contains returns true iff the set contains field.
func (fs FieldSet) Contains(field string) bool {
	_, found := fs[field]
	return found
}

// FieldError indicates an error setting a field in destination.
type FieldError struct {
	Message string
	Field   string
	Inner   error
}

func (e *FieldError) Error() string {
	if e.Inner == nil {
		return fmt.Sprintf(
			"invalid field %q in %s",
			e.Field, e.Message,
		)
	}
	return fmt.Sprintf(
		"invalid field %q in %s: %v",
		e.Field, e.Message, e.Inner,
	)
}

func (e *FieldError) Unwrap() error {
	return e.Inner
}

// WrapFieldError returns a *FieldError that wraps err.
func WrapFieldError(message, field string, err error) error {
	return &FieldError{
		Message: message,
		Field:   field,
		Inner:   err,
	}
}

// FieldErrorf returns a *FieldError that wraps fmt.Errorf(format, a...).
func FieldErrorf(message, field, format string, a ...interface{}) error {
	return WrapFieldError(message, field, fmt.Errorf(format, a...))
}

func depth(path string) int {
	return strings.Count(path, ".") + 1
}

// TopLevelPaths returns the paths that do not have sub-paths.
func TopLevelPaths(paths []string) []string {
	out := make([]string, 0, len(paths))
	for _, path := range paths {
		if depth(path) == 1 {
			out = append(out, path)
		}
	}
	return out
}

// SubPaths returns the paths that have sub-paths.
func SubPaths(paths []string) []string {
	out := make([]string, 0, len(paths))
	for _, path := range paths {
		if depth(path) > 1 {
			out = append(out, path)
		}
	}
	return out
}

// SubPathsOf returns the sub-paths of all paths that start with topLevelField.
func SubPathsOf(paths []string, topLevelField string) []string {
	prefix := topLevelField + "."
	out := make([]string, 0, len(paths))
	for _, path := range paths {
		if strings.HasPrefix(path, prefix) {
			out = append(out, strings.TrimPrefix(path, prefix))
		}
	}
	return out
}

// TopLevelField returns the top-level field of the path (meaning: everything before the first period).
func TopLevelField(path string) string {
	sep := "."
	if i := strings.Index(path, sep); i >= 0 {
		return path[:i]
	}
	return path
}

// PrefixPaths adds the given prefix to the paths.
func PrefixPaths(paths []string, prefix string) []string {
	out := make([]string, len(paths))
	for i, v := range paths {
		out[i] = prefix + "." + v
	}
	return out
}

// Flatten flattens a list of string slices.
func Flatten(paths [][]string) []string {
	var outLen int
	for _, paths := range paths {
		outLen += len(paths)
	}
	out := make([]string, 0, outLen)
	for _, paths := range paths {
		out = append(out, paths...)
	}
	return out
}
