// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package fieldsetterplugin

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
	Field       string
	Destination interface{}
	Inner       error
}

func (e *FieldError) Error() string {
	if e.Inner == nil {
		return fmt.Sprintf(
			"can not set field %q in %T",
			e.Field, e.Destination,
		)
	}
	return fmt.Sprintf(
		"can not set field %q in %T: %v",
		e.Field, e.Destination, e.Inner,
	)
}

func (e *FieldError) Unwrap() error {
	return e.Inner
}

// WrapFieldError returns a *FieldError that wraps err.
func WrapFieldError(dst interface{}, field string, err error) error {
	return &FieldError{
		Destination: dst,
		Field:       field,
		Inner:       err,
	}
}

// FieldErrorf returns a *FieldError that wraps fmt.Errorf(format, a...).
func FieldErrorf(dst interface{}, field, format string, a ...interface{}) error {
	return WrapFieldError(dst, field, fmt.Errorf(format, a...))
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
	field, _, _ := strings.Cut(path, ".")
	return field
}
