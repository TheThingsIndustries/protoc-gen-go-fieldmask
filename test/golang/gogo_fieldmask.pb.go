// Code generated by protoc-gen-go-fieldmask. DO NOT EDIT.
// versions:
// - protoc-gen-go-fieldmask v0.0.0-dev
// - protoc             v3.20.1
// source: gogo.proto

package test

import (
	fmt "fmt"
	fieldmaskplugin "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/fieldmaskplugin"
)

// SetFields sets the given fields from src into x.
func (x *MessageWithGoGoOptions) SetFields(src *MessageWithGoGoOptions, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithGoGoOptions")
	case src == nil:
		src = &MessageWithGoGoOptions{}
	}
	fset := make(fieldmaskplugin.FieldSet, 8)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldmaskplugin.FieldErrorf("MessageWithGoGoOptions", field, "unknown field")
		case "eui_with_custom_name", "euiWithCustomName":
			x.EuiWithCustomName = src.EuiWithCustomName
		case "eui_with_custom_name_and_type", "euiWithCustomNameAndType":
			x.EuiWithCustomNameAndType = src.EuiWithCustomNameAndType
		case "non_nullable_eui_with_custom_name_and_type", "nonNullableEuiWithCustomNameAndType":
			x.NonNullableEuiWithCustomNameAndType = src.NonNullableEuiWithCustomNameAndType
		case "euis_with_custom_name_and_type", "euisWithCustomNameAndType":
			x.EuisWithCustomNameAndType = src.EuisWithCustomNameAndType
		case "duration":
			x.Duration = src.Duration
		case "non_nullable_duration", "nonNullableDuration":
			x.NonNullableDuration = src.NonNullableDuration
		case "timestamp":
			x.Timestamp = src.Timestamp
		case "non_nullable_timestamp", "nonNullableTimestamp":
			x.NonNullableTimestamp = src.NonNullableTimestamp
		}
		fset.Add(field)
	}
	return nil
}

// SetFields sets the given fields from src into x.
func (x *MessageWithNullable) SetFields(src *MessageWithNullable, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithNullable")
	case src == nil:
		src = &MessageWithNullable{}
	}
	fset := make(fieldmaskplugin.FieldSet, 2)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldmaskplugin.FieldErrorf("MessageWithNullable", field, "unknown field")
		case "sub":
			x.Sub = src.Sub
		case "subs":
			x.Subs = src.Subs
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
			return fieldmaskplugin.FieldErrorf("MessageWithNullable", field, "unknown field")
		case "sub":
			if x.Sub == nil && src.Sub != nil {
				var v SubMessage
				x.Sub = &v
			}
			if err := x.Sub.SetFields(src.Sub, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldmaskplugin.WrapFieldError("MessageWithNullable", field, err)
			}
		}
		fset.Add(topLevelField)
	}
	return nil
}

// SetFields sets the given fields from src into x.
func (x *MessageWithEmbedded) SetFields(src *MessageWithEmbedded, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithEmbedded")
	case src == nil:
		src = &MessageWithEmbedded{}
	}
	fset := make(fieldmaskplugin.FieldSet, 1)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldmaskplugin.FieldErrorf("MessageWithEmbedded", field, "unknown field")
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
			return fieldmaskplugin.FieldErrorf("MessageWithEmbedded", field, "unknown field")
		case "sub":
			if x.Sub == nil && src.Sub != nil {
				var v SubMessage
				x.Sub = &v
			}
			if err := x.Sub.SetFields(src.Sub, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldmaskplugin.WrapFieldError("MessageWithEmbedded", field, err)
			}
		}
		fset.Add(topLevelField)
	}
	return nil
}

// SetFields sets the given fields from src into x.
func (x *MessageWithNullableEmbedded) SetFields(src *MessageWithNullableEmbedded, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithNullableEmbedded")
	case src == nil:
		src = &MessageWithNullableEmbedded{}
	}
	fset := make(fieldmaskplugin.FieldSet, 1)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldmaskplugin.FieldErrorf("MessageWithNullableEmbedded", field, "unknown field")
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
			return fieldmaskplugin.FieldErrorf("MessageWithNullableEmbedded", field, "unknown field")
		case "sub":
			if x.Sub == nil && src.Sub != nil {
				var v SubMessage
				x.Sub = &v
			}
			if err := x.Sub.SetFields(src.Sub, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldmaskplugin.WrapFieldError("MessageWithNullableEmbedded", field, err)
			}
		}
		fset.Add(topLevelField)
	}
	return nil
}
