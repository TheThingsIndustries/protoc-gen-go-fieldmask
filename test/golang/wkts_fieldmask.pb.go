// Code generated by protoc-gen-go-fieldmask. DO NOT EDIT.
// versions:
// - protoc-gen-go-fieldmask v0.0.0-dev
// - protoc             v3.20.1
// source: wkts.proto

package test

import (
	fmt "fmt"
	fieldmaskplugin "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/fieldmaskplugin"
)

// FieldPaths returns the field paths up to the given maximum depth.
func (x *MessageWithWKTs) FieldPaths(maxDepth int) []string {
	if maxDepth == 0 {
		return nil
	}
	return []string{
		"double_value",
		"double_values",
		"float_value",
		"float_values",
		"int32_value",
		"int32_values",
		"int64_value",
		"int64_values",
		"uint32_value",
		"uint32_values",
		"uint64_value",
		"uint64_values",
		"bool_value",
		"bool_values",
		"string_value",
		"string_values",
		"bytes_value",
		"bytes_values",
		"empty_value",
		"empty_values",
		"timestamp_value",
		"timestamp_values",
		"duration_value",
		"duration_values",
		"field_mask_value",
		"field_mask_values",
		"value_value",
		"value_values",
		"list_value_value",
		"list_value_values",
		"struct_value",
		"struct_values",
		"any_value",
		"any_values",
	}
}

// SetFields sets the given fields from src into x.
func (x *MessageWithWKTs) SetFields(src *MessageWithWKTs, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithWKTs")
	case src == nil:
		src = &MessageWithWKTs{}
	}
	fset := make(fieldmaskplugin.FieldSet, 34)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldmaskplugin.FieldErrorf("MessageWithWKTs", field, "unknown field")
		case "double_value", "doubleValue":
			x.DoubleValue = src.DoubleValue
		case "double_values", "doubleValues":
			x.DoubleValues = src.DoubleValues
		case "float_value", "floatValue":
			x.FloatValue = src.FloatValue
		case "float_values", "floatValues":
			x.FloatValues = src.FloatValues
		case "int32_value", "int32Value":
			x.Int32Value = src.Int32Value
		case "int32_values", "int32Values":
			x.Int32Values = src.Int32Values
		case "int64_value", "int64Value":
			x.Int64Value = src.Int64Value
		case "int64_values", "int64Values":
			x.Int64Values = src.Int64Values
		case "uint32_value", "uint32Value":
			x.Uint32Value = src.Uint32Value
		case "uint32_values", "uint32Values":
			x.Uint32Values = src.Uint32Values
		case "uint64_value", "uint64Value":
			x.Uint64Value = src.Uint64Value
		case "uint64_values", "uint64Values":
			x.Uint64Values = src.Uint64Values
		case "bool_value", "boolValue":
			x.BoolValue = src.BoolValue
		case "bool_values", "boolValues":
			x.BoolValues = src.BoolValues
		case "string_value", "stringValue":
			x.StringValue = src.StringValue
		case "string_values", "stringValues":
			x.StringValues = src.StringValues
		case "bytes_value", "bytesValue":
			x.BytesValue = src.BytesValue
		case "bytes_values", "bytesValues":
			x.BytesValues = src.BytesValues
		case "empty_value", "emptyValue":
			x.EmptyValue = src.EmptyValue
		case "empty_values", "emptyValues":
			x.EmptyValues = src.EmptyValues
		case "timestamp_value", "timestampValue":
			x.TimestampValue = src.TimestampValue
		case "timestamp_values", "timestampValues":
			x.TimestampValues = src.TimestampValues
		case "duration_value", "durationValue":
			x.DurationValue = src.DurationValue
		case "duration_values", "durationValues":
			x.DurationValues = src.DurationValues
		case "field_mask_value", "fieldMaskValue":
			x.FieldMaskValue = src.FieldMaskValue
		case "field_mask_values", "fieldMaskValues":
			x.FieldMaskValues = src.FieldMaskValues
		case "value_value", "valueValue":
			x.ValueValue = src.ValueValue
		case "value_values", "valueValues":
			x.ValueValues = src.ValueValues
		case "list_value_value", "listValueValue":
			x.ListValueValue = src.ListValueValue
		case "list_value_values", "listValueValues":
			x.ListValueValues = src.ListValueValues
		case "struct_value", "structValue":
			x.StructValue = src.StructValue
		case "struct_values", "structValues":
			x.StructValues = src.StructValues
		case "any_value", "anyValue":
			x.AnyValue = src.AnyValue
		case "any_values", "anyValues":
			x.AnyValues = src.AnyValues
		}
		fset.Add(field)
	}
	return nil
}

// FieldPaths returns the field paths up to the given maximum depth.
func (x *MessageWithOneofWKTs) FieldPaths(maxDepth int) []string {
	if maxDepth == 0 {
		return nil
	}
	return []string{
		"double_value",
		"float_value",
		"int32_value",
		"int64_value",
		"uint32_value",
		"uint64_value",
		"bool_value",
		"string_value",
		"bytes_value",
		"empty_value",
		"timestamp_value",
		"duration_value",
		"field_mask_value",
		"value_value",
		"list_value_value",
		"struct_value",
		"any_value",
	}
}

// SetFields sets the given fields from src into x.
func (x *MessageWithOneofWKTs) SetFields(src *MessageWithOneofWKTs, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithOneofWKTs")
	case src == nil:
		src = &MessageWithOneofWKTs{}
	}
	fset := make(fieldmaskplugin.FieldSet, 18)
	for _, field := range fieldmaskplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "unknown field")
		case "double_value", "doubleValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_DoubleValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "float_value", "floatValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_FloatValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "int32_value", "int32Value":
			ov, ok := src.Value.(*MessageWithOneofWKTs_Int32Value)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "int64_value", "int64Value":
			ov, ok := src.Value.(*MessageWithOneofWKTs_Int64Value)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "uint32_value", "uint32Value":
			ov, ok := src.Value.(*MessageWithOneofWKTs_Uint32Value)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "uint64_value", "uint64Value":
			ov, ok := src.Value.(*MessageWithOneofWKTs_Uint64Value)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "bool_value", "boolValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_BoolValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "string_value", "stringValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_StringValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "bytes_value", "bytesValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_BytesValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "empty_value", "emptyValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_EmptyValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "timestamp_value", "timestampValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_TimestampValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "duration_value", "durationValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_DurationValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "field_mask_value", "fieldMaskValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_FieldMaskValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "value_value", "valueValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_ValueValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "list_value_value", "listValueValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_ListValueValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "struct_value", "structValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_StructValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "any_value", "anyValue":
			ov, ok := src.Value.(*MessageWithOneofWKTs_AnyValue)
			if !ok {
				return fieldmaskplugin.FieldErrorf("MessageWithOneofWKTs", field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "value":
			x.Value = src.Value
		}
		fset.Add(field)
	}
	return nil
}
