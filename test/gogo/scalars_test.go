// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package test_test

import (
	"testing"

	. "github.com/TheThingsIndustries/protoc-gen-go-field-setters/test/gogo"
)

var fullMessageWithScalars = &MessageWithScalars{
	DoubleValue:    12.34,
	DoubleValues:   []float64{12.34, 56.78},
	FloatValue:     12.34,
	FloatValues:    []float32{12.34, 56.78},
	Int32Value:     -42,
	Int32Values:    []int32{1, 2, -42},
	Int64Value:     -42,
	Int64Values:    []int64{1, 2, -42},
	Uint32Value:    42,
	Uint32Values:   []uint32{1, 2, 42},
	Uint64Value:    42,
	Uint64Values:   []uint64{1, 2, 42},
	Sint32Value:    -42,
	Sint32Values:   []int32{1, 2, -42},
	Sint64Value:    -42,
	Sint64Values:   []int64{1, 2, -42},
	Fixed32Value:   42,
	Fixed32Values:  []uint32{1, 2, 42},
	Fixed64Value:   42,
	Fixed64Values:  []uint64{1, 2, 42},
	Sfixed32Value:  -42,
	Sfixed32Values: []int32{1, 2, -42},
	Sfixed64Value:  -42,
	Sfixed64Values: []int64{1, 2, -42},
	BoolValue:      true,
	BoolValues:     []bool{true, false},
	StringValue:    "foo",
	StringValues:   []string{"foo", "bar"},
	BytesValue:     []byte("foo"),
	BytesValues:    [][]byte{[]byte("foo"), []byte("bar")},
}

func TestMessageWithScalars(t *testing.T) {
	var dst MessageWithScalars

	err := dst.SetFields(
		fullMessageWithScalars,
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
		"sint32_value",
		"sint32_values",
		"sint64_value",
		"sint64_values",
		"fixed32_value",
		"fixed32_values",
		"fixed64_value",
		"fixed64_values",
		"sfixed32_value",
		"sfixed32_values",
		"sfixed64_value",
		"sfixed64_values",
		"bool_value",
		"bool_values",
		"string_value",
		"string_values",
		"bytes_value",
		"bytes_values",
	)
	if err != nil {
		t.Fatal(err)
	}

	expectMessage(t, fullMessageWithScalars, &dst)
}

var testMessagesWithOneofScalars = []struct {
	name string
	msg  *MessageWithOneofScalars
	path string
}{
	{
		name: "empty",
		msg:  &MessageWithOneofScalars{},
		path: "value",
	},
	{
		name: "double_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_DoubleValue{DoubleValue: 12.34},
		},
		path: "double_value",
	},
	{
		name: "float_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_FloatValue{FloatValue: 12.34},
		},
		path: "float_value",
	},
	{
		name: "int32_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Int32Value{Int32Value: -42},
		},
		path: "int32_value",
	},
	{
		name: "int64_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Int64Value{Int64Value: -42},
		},
		path: "int64_value",
	},
	{
		name: "uint32_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Uint32Value{Uint32Value: 42},
		},
		path: "uint32_value",
	},
	{
		name: "uint64_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Uint64Value{Uint64Value: 42},
		},
		path: "uint64_value",
	},
	{
		name: "sint32_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Sint32Value{Sint32Value: -42},
		},
		path: "sint32_value",
	},
	{
		name: "sint64_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Sint64Value{Sint64Value: -42},
		},
		path: "sint64_value",
	},
	{
		name: "fixed32_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Fixed32Value{Fixed32Value: 42},
		},
		path: "fixed32_value",
	},
	{
		name: "fixed64_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Fixed64Value{Fixed64Value: 42},
		},
		path: "fixed64_value",
	},

	{
		name: "sfixed32_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Sfixed32Value{Sfixed32Value: -42},
		},
		path: "sfixed32_value",
	},
	{
		name: "sfixed64_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Sfixed64Value{Sfixed64Value: -42},
		},
		path: "sfixed64_value",
	},
	{
		name: "bool_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_BoolValue{BoolValue: true},
		},
		path: "bool_value",
	},
	{
		name: "string_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_StringValue{StringValue: "foo"},
		},
		path: "string_value",
	},
	{
		name: "bytes_value",
		msg: &MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_BytesValue{BytesValue: []byte("foo")},
		},
		path: "bytes_value",
	},
}

func TestMessageWithOneofScalars(t *testing.T) {
	for _, tt := range testMessagesWithOneofScalars {
		t.Run(tt.name, func(t *testing.T) {
			var dst MessageWithOneofScalars

			err := dst.SetFields(tt.msg, tt.path)
			if err != nil {
				t.Fatal(err)
			}

			expectMessage(t, tt.msg, &dst)
		})
	}
}

var fullMessageWithScalarMaps = &MessageWithScalarMaps{
	StringDoubleMap:   map[string]float64{"value": -42},
	StringFloatMap:    map[string]float32{"value": -42},
	StringInt32Map:    map[string]int32{"value": -42},
	Int32StringMap:    map[int32]string{-42: "answer"},
	StringInt64Map:    map[string]int64{"value": -42},
	Int64StringMap:    map[int64]string{-42: "answer"},
	StringUint32Map:   map[string]uint32{"value": 42},
	Uint32StringMap:   map[uint32]string{42: "answer"},
	StringUint64Map:   map[string]uint64{"value": 42},
	Uint64StringMap:   map[uint64]string{42: "answer"},
	StringSint32Map:   map[string]int32{"value": -42},
	Sint32StringMap:   map[int32]string{-42: "answer"},
	StringSint64Map:   map[string]int64{"value": -42},
	Sint64StringMap:   map[int64]string{-42: "answer"},
	StringFixed32Map:  map[string]uint32{"value": 42},
	Fixed32StringMap:  map[uint32]string{42: "answer"},
	StringFixed64Map:  map[string]uint64{"value": 42},
	Fixed64StringMap:  map[uint64]string{42: "answer"},
	StringSfixed32Map: map[string]int32{"value": -42},
	Sfixed32StringMap: map[int32]string{-42: "answer"},
	StringSfixed64Map: map[string]int64{"value": -42},
	Sfixed64StringMap: map[int64]string{-42: "answer"},
	StringBoolMap:     map[string]bool{"yes": true},
	BoolStringMap:     map[bool]string{true: "yes"},
	StringStringMap:   map[string]string{"value": "foo"},
	StringBytesMap:    map[string][]byte{"value": []byte("foo")},
}

func TestMessageWithScalarMaps(t *testing.T) {
	var dst MessageWithScalarMaps

	err := dst.SetFields(
		fullMessageWithScalarMaps,
		"string_double_map",
		"string_float_map",
		"string_int32_map",
		"int32_string_map",
		"string_int64_map",
		"int64_string_map",
		"string_uint32_map",
		"uint32_string_map",
		"string_uint64_map",
		"uint64_string_map",
		"string_sint32_map",
		"sint32_string_map",
		"string_sint64_map",
		"sint64_string_map",
		"string_fixed32_map",
		"fixed32_string_map",
		"string_fixed64_map",
		"fixed64_string_map",
		"string_sfixed32_map",
		"sfixed32_string_map",
		"string_sfixed64_map",
		"sfixed64_string_map",
		"string_bool_map",
		"bool_string_map",
		"string_string_map",
		"string_bytes_map",
	)
	if err != nil {
		t.Fatal(err)
	}

	expectMessage(t, fullMessageWithScalarMaps, &dst)
}
