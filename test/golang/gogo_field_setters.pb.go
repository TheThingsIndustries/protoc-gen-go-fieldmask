// Code generated by protoc-gen-go-field-setters. DO NOT EDIT.
// versions:
// - protoc-gen-go-field-setters v0.0.0-dev
// - protoc             v3.20.1
// source: gogo.proto

package test

import (
	fmt "fmt"
	fieldsetterplugin "github.com/TheThingsIndustries/protoc-gen-go-field-setters/fieldsetterplugin"
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
	fset := make(fieldsetterplugin.FieldSet)
	for _, field := range fieldsetterplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldsetterplugin.FieldErrorf(x, field, "unknown field")
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
	fset := make(fieldsetterplugin.FieldSet)
	for _, field := range fieldsetterplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldsetterplugin.FieldErrorf(x, field, "unknown field")
		case "sub":
			x.Sub = src.Sub
		case "subs":
			x.Subs = src.Subs
		}
		fset.Add(field)
	}
	for _, field := range fieldsetterplugin.SubPaths(paths) {
		topLevelField := fieldsetterplugin.TopLevelField(field)
		if fset.Contains(topLevelField) {
			continue
		}
		switch topLevelField {
		default:
			return fieldsetterplugin.FieldErrorf(x, field, "unknown field")
		case "sub":
			if x.Sub == nil && src.Sub != nil {
				var v SubMessage
				x.Sub = &v
			}
			if err := x.Sub.SetFields(src.Sub, fieldsetterplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldsetterplugin.WrapFieldError(x, field, err)
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
	fset := make(fieldsetterplugin.FieldSet)
	for _, field := range fieldsetterplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldsetterplugin.FieldErrorf(x, field, "unknown field")
		case "sub":
			x.Sub = src.Sub
		}
		fset.Add(field)
	}
	for _, field := range fieldsetterplugin.SubPaths(paths) {
		topLevelField := fieldsetterplugin.TopLevelField(field)
		if fset.Contains(topLevelField) {
			continue
		}
		switch topLevelField {
		default:
			return fieldsetterplugin.FieldErrorf(x, field, "unknown field")
		case "sub":
			if x.Sub == nil && src.Sub != nil {
				var v SubMessage
				x.Sub = &v
			}
			if err := x.Sub.SetFields(src.Sub, fieldsetterplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldsetterplugin.WrapFieldError(x, field, err)
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
	fset := make(fieldsetterplugin.FieldSet)
	for _, field := range fieldsetterplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldsetterplugin.FieldErrorf(x, field, "unknown field")
		case "sub":
			x.Sub = src.Sub
		}
		fset.Add(field)
	}
	for _, field := range fieldsetterplugin.SubPaths(paths) {
		topLevelField := fieldsetterplugin.TopLevelField(field)
		if fset.Contains(topLevelField) {
			continue
		}
		switch topLevelField {
		default:
			return fieldsetterplugin.FieldErrorf(x, field, "unknown field")
		case "sub":
			if x.Sub == nil && src.Sub != nil {
				var v SubMessage
				x.Sub = &v
			}
			if err := x.Sub.SetFields(src.Sub, fieldsetterplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldsetterplugin.WrapFieldError(x, field, err)
			}
		}
		fset.Add(topLevelField)
	}
	return nil
}
