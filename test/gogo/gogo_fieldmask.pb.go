// Code generated by protoc-gen-go-fieldmask. DO NOT EDIT.
// versions:
// - protoc-gen-go-fieldmask v0.0.0-dev
// - protoc             v3.9.1
// source: gogo.proto

package test

import (
	fmt "fmt"
	fieldmaskplugin "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/fieldmaskplugin"
)

// FieldPaths returns the field paths up to the given maximum depth.
func (x *MessageWithGoGoOptions) FieldPaths(maxDepth int) []string {
	if maxDepth == 0 {
		return nil
	}
	return []string{
		"eui_with_custom_name",
		"eui_with_custom_name_and_type",
		"non_nullable_eui_with_custom_name_and_type",
		"euis_with_custom_name_and_type",
		"duration",
		"non_nullable_duration",
		"timestamp",
		"non_nullable_timestamp",
	}
}

// NormalizeFieldPaths normalizes the field paths.
func (x *MessageWithGoGoOptions) NormalizeFieldPaths(paths ...string) ([]string, error) {
	var (
		normalizedPaths []string
		fset            = make(fieldmaskplugin.FieldSet, 8)
	)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return nil, fieldmaskplugin.FieldErrorf("MessageWithGoGoOptions", field, "unknown field")
		case "eui_with_custom_name", "euiWithCustomName":
			normalizedPaths = append(normalizedPaths, "eui_with_custom_name")
		case "eui_with_custom_name_and_type", "euiWithCustomNameAndType":
			normalizedPaths = append(normalizedPaths, "eui_with_custom_name_and_type")
		case "non_nullable_eui_with_custom_name_and_type", "nonNullableEuiWithCustomNameAndType":
			normalizedPaths = append(normalizedPaths, "non_nullable_eui_with_custom_name_and_type")
		case "euis_with_custom_name_and_type", "euisWithCustomNameAndType":
			normalizedPaths = append(normalizedPaths, "euis_with_custom_name_and_type")
		case "duration":
			normalizedPaths = append(normalizedPaths, "duration")
		case "non_nullable_duration", "nonNullableDuration":
			normalizedPaths = append(normalizedPaths, "non_nullable_duration")
		case "timestamp":
			normalizedPaths = append(normalizedPaths, "timestamp")
		case "non_nullable_timestamp", "nonNullableTimestamp":
			normalizedPaths = append(normalizedPaths, "non_nullable_timestamp")
		}
		fset.Add(field)
	}
	return normalizedPaths, nil
}

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
			x.EUIWithCustomName = src.EUIWithCustomName
		case "eui_with_custom_name_and_type", "euiWithCustomNameAndType":
			x.EUIWithCustomNameAndType = src.EUIWithCustomNameAndType
		case "non_nullable_eui_with_custom_name_and_type", "nonNullableEuiWithCustomNameAndType":
			x.NonNullableEUIWithCustomNameAndType = src.NonNullableEUIWithCustomNameAndType
		case "euis_with_custom_name_and_type", "euisWithCustomNameAndType":
			x.EUIsWithCustomNameAndType = src.EUIsWithCustomNameAndType
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

// FieldPaths returns the field paths up to the given maximum depth.
func (x *MessageWithNullable) FieldPaths(maxDepth int) []string {
	if maxDepth == 0 {
		return nil
	}
	if maxDepth > 1 {
		return fieldmaskplugin.Flatten([][]string{
			{"sub"},
			fieldmaskplugin.PrefixPaths(((*SubMessage)(nil)).FieldPaths(maxDepth-1), "sub"),
			{"subs"},
		})
	}
	return []string{
		"sub",
		"subs",
	}
}

// NormalizeFieldPaths normalizes the field paths.
func (x *MessageWithNullable) NormalizeFieldPaths(paths ...string) ([]string, error) {
	var (
		normalizedPaths []string
		fset            = make(fieldmaskplugin.FieldSet, 2)
	)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return nil, fieldmaskplugin.FieldErrorf("MessageWithNullable", field, "unknown field")
		case "sub":
			normalizedPaths = append(normalizedPaths, "sub")
		case "subs":
			normalizedPaths = append(normalizedPaths, "subs")
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
			return nil, fieldmaskplugin.FieldErrorf("MessageWithNullable", field, "unknown field")
		case "sub":
			normalizedSubFields, err := ((*SubMessage)(nil)).NormalizeFieldPaths(fieldmaskplugin.SubPathsOf(paths, topLevelField)...)
			if err != nil {
				return nil, err
			}
			normalizedPaths = append(normalizedPaths, fieldmaskplugin.PrefixPaths(normalizedSubFields, topLevelField)...)
		}
		fset.Add(topLevelField)
	}
	return normalizedPaths, nil
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
			if err := x.Sub.SetFields(&src.Sub, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldmaskplugin.WrapFieldError("MessageWithNullable", field, err)
			}
		}
		fset.Add(topLevelField)
	}
	return nil
}

// FieldPaths returns the field paths up to the given maximum depth.
func (x *MessageWithEmbedded) FieldPaths(maxDepth int) []string {
	if maxDepth == 0 {
		return nil
	}
	if maxDepth > 1 {
		return fieldmaskplugin.Flatten([][]string{
			{"sub"},
			fieldmaskplugin.PrefixPaths(((*SubMessage)(nil)).FieldPaths(maxDepth-1), "sub"),
		})
	}
	return []string{
		"sub",
	}
}

// NormalizeFieldPaths normalizes the field paths.
func (x *MessageWithEmbedded) NormalizeFieldPaths(paths ...string) ([]string, error) {
	var (
		normalizedPaths []string
		fset            = make(fieldmaskplugin.FieldSet, 1)
	)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return nil, fieldmaskplugin.FieldErrorf("MessageWithEmbedded", field, "unknown field")
		case "sub":
			normalizedPaths = append(normalizedPaths, "sub")
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
			return nil, fieldmaskplugin.FieldErrorf("MessageWithEmbedded", field, "unknown field")
		case "sub":
			normalizedSubFields, err := ((*SubMessage)(nil)).NormalizeFieldPaths(fieldmaskplugin.SubPathsOf(paths, topLevelField)...)
			if err != nil {
				return nil, err
			}
			normalizedPaths = append(normalizedPaths, fieldmaskplugin.PrefixPaths(normalizedSubFields, topLevelField)...)
		}
		fset.Add(topLevelField)
	}
	return normalizedPaths, nil
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
			x.SubMessage = src.SubMessage
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
			if x.SubMessage == nil && src.SubMessage != nil {
				var v SubMessage
				x.SubMessage = &v
			}
			if err := x.SubMessage.SetFields(src.SubMessage, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldmaskplugin.WrapFieldError("MessageWithEmbedded", field, err)
			}
		}
		fset.Add(topLevelField)
	}
	return nil
}

// FieldPaths returns the field paths up to the given maximum depth.
func (x *MessageWithNullableEmbedded) FieldPaths(maxDepth int) []string {
	if maxDepth == 0 {
		return nil
	}
	if maxDepth > 1 {
		return fieldmaskplugin.Flatten([][]string{
			{"sub"},
			fieldmaskplugin.PrefixPaths(((*SubMessage)(nil)).FieldPaths(maxDepth-1), "sub"),
		})
	}
	return []string{
		"sub",
	}
}

// NormalizeFieldPaths normalizes the field paths.
func (x *MessageWithNullableEmbedded) NormalizeFieldPaths(paths ...string) ([]string, error) {
	var (
		normalizedPaths []string
		fset            = make(fieldmaskplugin.FieldSet, 1)
	)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return nil, fieldmaskplugin.FieldErrorf("MessageWithNullableEmbedded", field, "unknown field")
		case "sub":
			normalizedPaths = append(normalizedPaths, "sub")
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
			return nil, fieldmaskplugin.FieldErrorf("MessageWithNullableEmbedded", field, "unknown field")
		case "sub":
			normalizedSubFields, err := ((*SubMessage)(nil)).NormalizeFieldPaths(fieldmaskplugin.SubPathsOf(paths, topLevelField)...)
			if err != nil {
				return nil, err
			}
			normalizedPaths = append(normalizedPaths, fieldmaskplugin.PrefixPaths(normalizedSubFields, topLevelField)...)
		}
		fset.Add(topLevelField)
	}
	return normalizedPaths, nil
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
			x.SubMessage = src.SubMessage
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
			if err := x.SubMessage.SetFields(&src.SubMessage, fieldmaskplugin.SubPathsOf(paths, topLevelField)...); err != nil {
				return fieldmaskplugin.WrapFieldError("MessageWithNullableEmbedded", field, err)
			}
		}
		fset.Add(topLevelField)
	}
	return nil
}
