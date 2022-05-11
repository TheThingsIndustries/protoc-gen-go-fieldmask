// Code generated by protoc-gen-go-fieldmask. DO NOT EDIT.
// versions:
// - protoc-gen-go-fieldmask v0.0.0-dev
// - protoc             v3.9.1
// source: messages.proto

package test

import (
	fmt "fmt"
	fieldmaskplugin "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/fieldmaskplugin"
)

// FieldPaths returns the field paths up to the given maximum depth.
func (x *MessageWithoutFieldSetter) FieldPaths(maxDepth int) []string {
	if maxDepth == 0 {
		return nil
	}
	return []string{
		"message",
	}
}

// FieldPaths returns the field paths up to the given maximum depth.
func (x *SubMessage) FieldPaths(maxDepth int) []string {
	if maxDepth == 0 {
		return nil
	}
	return []string{
		"foo",
		"bar",
	}
}

// SetFields sets the given fields from src into x.
func (x *SubMessage) SetFields(src *SubMessage, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil SubMessage")
	case src == nil:
		src = &SubMessage{}
	}
	fset := make(fieldmaskplugin.FieldSet, 2)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldmaskplugin.FieldErrorf("SubMessage", field, "unknown field")
		case "foo":
			x.Foo = src.Foo
		case "bar":
			x.Bar = src.Bar
		}
		fset.Add(field)
	}
	return nil
}

// FieldPaths returns the field paths up to the given maximum depth.
func (x *MessageWithSubMessages) FieldPaths(maxDepth int) []string {
	if maxDepth == 0 {
		return nil
	}
	if maxDepth > 1 {
		return fieldmaskplugin.Flatten([][]string{
			{"a"},
			fieldmaskplugin.PrefixPaths(((*SubMessage)(nil)).FieldPaths(maxDepth-1), "a"),
			{"b"},
			fieldmaskplugin.PrefixPaths(((*SubMessage)(nil)).FieldPaths(maxDepth-1), "b"),
		})
	}
	return []string{
		"a",
		"b",
	}
}

// SetFields sets the given fields from src into x.
func (x *MessageWithSubMessages) SetFields(src *MessageWithSubMessages, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithSubMessages")
	case src == nil:
		src = &MessageWithSubMessages{}
	}
	fset := make(fieldmaskplugin.FieldSet, 2)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldmaskplugin.FieldErrorf("MessageWithSubMessages", field, "unknown field")
		case "a":
			x.A = src.A
		case "b":
			x.B = src.B
		}
		fset.Add(field)
	}
	for _, field := range fieldmaskplugin.SubPaths(paths) {
		topLevelField := fieldmaskplugin.TopLevelField(field)
		if fset.Contains(topLevelField) {
			continue
		}
		switch topLevelField {
		default:
			return fieldmaskplugin.FieldErrorf("MessageWithSubMessages", field, "unknown field")
		case "a":
			if x.A == nil && src.A != nil {
				var v SubMessage
				x.A = &v
			}
			if err := x.A.SetFields(src.A, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldmaskplugin.WrapFieldError("MessageWithSubMessages", field, err)
			}
		case "b":
			if x.B == nil && src.B != nil {
				var v SubMessage
				x.B = &v
			}
			if err := x.B.SetFields(src.B, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldmaskplugin.WrapFieldError("MessageWithSubMessages", field, err)
			}
		}
		fset.Add(topLevelField)
	}
	return nil
}

// FieldPaths returns the field paths up to the given maximum depth.
func (x *MessageWithOneofSubMessages) FieldPaths(maxDepth int) []string {
	if maxDepth == 0 {
		return nil
	}
	if maxDepth > 1 {
		return fieldmaskplugin.Flatten([][]string{
			{"a"},
			fieldmaskplugin.PrefixPaths(((*SubMessage)(nil)).FieldPaths(maxDepth-1), "a"),
			{"b"},
			fieldmaskplugin.PrefixPaths(((*SubMessage)(nil)).FieldPaths(maxDepth-1), "b"),
		})
	}
	return []string{
		"a",
		"b",
	}
}

// SetFields sets the given fields from src into x.
func (x *MessageWithOneofSubMessages) SetFields(src *MessageWithOneofSubMessages, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithOneofSubMessages")
	case src == nil:
		src = &MessageWithOneofSubMessages{}
	}
	fset := make(fieldmaskplugin.FieldSet, 3)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldmaskplugin.FieldErrorf("MessageWithOneofSubMessages", field, "unknown field")
		case "a":
			ov, ok := src.Sub.(*MessageWithOneofSubMessages_A)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofSubMessages", field, "invalid Sub of type %T in source struct", src.Sub)
			}
			x.Sub = ov
		case "b":
			ov, ok := src.Sub.(*MessageWithOneofSubMessages_B)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofSubMessages", field, "invalid Sub of type %T in source struct", src.Sub)
			}
			x.Sub = ov
		case "sub":
			x.Sub = src.Sub
		}
		fset.Add(field)
	}
	for _, field := range fieldmaskplugin.SubPaths(paths) {
		topLevelField := fieldmaskplugin.TopLevelField(field)
		if fset.Contains(topLevelField) {
			continue
		}
		switch topLevelField {
		default:
			return fieldmaskplugin.FieldErrorf("MessageWithOneofSubMessages", field, "unknown field")
		case "a":
			if x.Sub == nil && src.Sub != nil {
				ov := &MessageWithOneofSubMessages_A{A: &SubMessage{}}
				x.Sub = ov
			}
			xOV, ok := x.Sub.(*MessageWithOneofSubMessages_A)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofSubMessages", field, "invalid Sub of type %T in destination struct", x.Sub)
			}
			srcOV, ok := src.Sub.(*MessageWithOneofSubMessages_A)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofSubMessages", field, "invalid Sub of type %T in source struct", src.Sub)
			}
			if err := xOV.A.SetFields(srcOV.A, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldmaskplugin.WrapFieldError("MessageWithOneofSubMessages", field, err)
			}
		case "b":
			if x.Sub == nil && src.Sub != nil {
				ov := &MessageWithOneofSubMessages_B{B: &SubMessage{}}
				x.Sub = ov
			}
			xOV, ok := x.Sub.(*MessageWithOneofSubMessages_B)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofSubMessages", field, "invalid Sub of type %T in destination struct", x.Sub)
			}
			srcOV, ok := src.Sub.(*MessageWithOneofSubMessages_B)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofSubMessages", field, "invalid Sub of type %T in source struct", src.Sub)
			}
			if err := xOV.B.SetFields(srcOV.B, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldmaskplugin.WrapFieldError("MessageWithOneofSubMessages", field, err)
			}
		case "sub":
			if err := x.SetFields(src, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return err
			}
		}
		fset.Add(topLevelField)
	}
	return nil
}
