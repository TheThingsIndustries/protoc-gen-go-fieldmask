// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package test_test

import (
	"testing"
	"time"

	. "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/test/gogo"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

var (
	testTime     = time.Date(2006, time.January, 2, 15, 4, 5, 123456789, time.FixedZone("07:00", 7*3600))
	testDuration = time.Hour + 2*time.Minute + 3*time.Second + 123456789
)

func mustTimestamp(t time.Time) *types.Timestamp {
	tmst, err := types.TimestampProto(t)
	if err != nil {
		panic(err)
	}
	return tmst
}

func mustDuration(t time.Duration) *types.Duration {
	return types.DurationProto(t)
}

func mustAny(pb proto.Message) *types.Any {
	any, err := types.MarshalAny(pb)
	if err != nil {
		panic(err)
	}
	return any
}

var fullMessageWithWKTs = &MessageWithWKTs{
	DoubleValue: &types.DoubleValue{Value: 12.34},
	DoubleValues: []*types.DoubleValue{
		{Value: 12.34},
		{Value: 56.78},
	},
	FloatValue: &types.FloatValue{Value: 12.34},
	FloatValues: []*types.FloatValue{
		{Value: 12.34},
		{Value: 56.78},
	},
	Int32Value: &types.Int32Value{Value: -42},
	Int32Values: []*types.Int32Value{
		{Value: 1},
		{Value: 2},
		{Value: -42},
	},
	Int64Value: &types.Int64Value{Value: -42},
	Int64Values: []*types.Int64Value{
		{Value: 1},
		{Value: 2},
		{Value: -42},
	},
	Uint32Value: &types.UInt32Value{Value: 42},
	Uint32Values: []*types.UInt32Value{
		{Value: 1},
		{Value: 2},
		{Value: 42},
	},
	Uint64Value: &types.UInt64Value{Value: 42},
	Uint64Values: []*types.UInt64Value{
		{Value: 1},
		{Value: 2},
		{Value: 42},
	},
	BoolValue: &types.BoolValue{Value: true},
	BoolValues: []*types.BoolValue{
		{Value: true},
		{Value: false},
	},
	StringValue: &types.StringValue{Value: "foo"},
	StringValues: []*types.StringValue{
		{Value: "foo"},
		{Value: "bar"},
	},
	BytesValue: &types.BytesValue{Value: []byte("foo")},
	BytesValues: []*types.BytesValue{
		{Value: []byte("foo")},
		{Value: []byte("bar")},
	},
	EmptyValue:     &types.Empty{},
	EmptyValues:    []*types.Empty{{}, {}},
	TimestampValue: mustTimestamp(testTime),
	TimestampValues: []*types.Timestamp{
		mustTimestamp(testTime),
		mustTimestamp(testTime.Truncate(10)),
		mustTimestamp(testTime.Truncate(100)),
		mustTimestamp(testTime.Truncate(1000)),
		mustTimestamp(testTime.Truncate(10000)),
		mustTimestamp(testTime.Truncate(100000)),
		mustTimestamp(testTime.Truncate(1000000)),
		mustTimestamp(testTime.Truncate(10000000)),
		mustTimestamp(testTime.Truncate(100000000)),
		mustTimestamp(testTime.Truncate(1000000000)),
	},
	DurationValue: mustDuration(testDuration),
	DurationValues: []*types.Duration{
		mustDuration(testDuration),
		mustDuration(testDuration.Truncate(10)),
		mustDuration(testDuration.Truncate(100)),
		mustDuration(testDuration.Truncate(1000)),
		mustDuration(testDuration.Truncate(10000)),
		mustDuration(testDuration.Truncate(100000)),
		mustDuration(testDuration.Truncate(1000000)),
		mustDuration(testDuration.Truncate(10000000)),
		mustDuration(testDuration.Truncate(100000000)),
		mustDuration(testDuration.Truncate(1000000000)),
	},
	FieldMaskValue: &types.FieldMask{Paths: []string{"foo.bar", "bar", "baz.qux"}},
	FieldMaskValues: []*types.FieldMask{
		{Paths: []string{"foo.bar", "bar", "baz.qux"}},
	},
	ValueValue: &types.Value{Kind: &types.Value_StringValue{StringValue: "foo"}},
	ValueValues: []*types.Value{
		{Kind: &types.Value_NullValue{}},
		{Kind: &types.Value_NumberValue{NumberValue: 12.34}},
		{Kind: &types.Value_StringValue{StringValue: "foo"}},
		{Kind: &types.Value_BoolValue{BoolValue: true}},
	},
	ListValueValue: &types.ListValue{
		Values: []*types.Value{
			{Kind: &types.Value_NullValue{}},
			{Kind: &types.Value_NumberValue{NumberValue: 12.34}},
			{Kind: &types.Value_StringValue{StringValue: "foo"}},
			{Kind: &types.Value_BoolValue{BoolValue: true}},
		},
	},
	ListValueValues: []*types.ListValue{
		{
			Values: []*types.Value{
				{Kind: &types.Value_NullValue{}},
				{Kind: &types.Value_NumberValue{NumberValue: 12.34}},
				{Kind: &types.Value_StringValue{StringValue: "foo"}},
				{Kind: &types.Value_BoolValue{BoolValue: true}},
			},
		},
	},
	StructValue: &types.Struct{
		Fields: map[string]*types.Value{
			"null":   {Kind: &types.Value_NullValue{}},
			"number": {Kind: &types.Value_NumberValue{NumberValue: 12.34}},
			"string": {Kind: &types.Value_StringValue{StringValue: "foo"}},
			"bool":   {Kind: &types.Value_BoolValue{BoolValue: true}},
		},
	},
	StructValues: []*types.Struct{
		{Fields: map[string]*types.Value{"null": {Kind: &types.Value_NullValue{}}}},
		{Fields: map[string]*types.Value{"number": {Kind: &types.Value_NumberValue{NumberValue: 12.34}}}},
		{Fields: map[string]*types.Value{"string": {Kind: &types.Value_StringValue{StringValue: "foo"}}}},
		{Fields: map[string]*types.Value{"bool": {Kind: &types.Value_BoolValue{BoolValue: true}}}},
	},
	AnyValue: mustAny(&SubMessage{Foo: "foo"}),
	AnyValues: []*types.Any{
		mustAny(&SubMessage{Foo: "foo"}),
		mustAny(mustTimestamp(testTime)),
		mustAny(mustDuration(testDuration)),
		mustAny(&types.FieldMask{Paths: []string{"foo.bar", "bar", "baz.qux"}}),
		mustAny(&types.Value{Kind: &types.Value_StringValue{StringValue: "foo"}}),
	},
}

func TestMessageWithWKTs(t *testing.T) {
	var dst MessageWithWKTs

	err := dst.SetFields(
		fullMessageWithWKTs,
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
	)
	if err != nil {
		t.Fatal(err)
	}

	expectMessage(t, fullMessageWithWKTs, &dst)
}

var testMessagesWithOneofWKTs = []struct {
	name string
	msg  *MessageWithOneofWKTs
	path string
}{
	{
		name: "empty",
		msg:  &MessageWithOneofWKTs{},
		path: "value",
	},
	{
		name: "double_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_DoubleValue{DoubleValue: &types.DoubleValue{Value: 12.34}},
		},
		path: "double_value",
	},
	{
		name: "float_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_FloatValue{FloatValue: &types.FloatValue{Value: 12.34}},
		},
		path: "float_value",
	},
	{
		name: "int32_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_Int32Value{Int32Value: &types.Int32Value{Value: -42}},
		},
		path: "int32_value",
	},
	{
		name: "int64_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_Int64Value{Int64Value: &types.Int64Value{Value: -42}},
		},
		path: "int64_value",
	},
	{
		name: "uint32_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_Uint32Value{Uint32Value: &types.UInt32Value{Value: 42}},
		},
		path: "uint32_value",
	},
	{
		name: "uint64_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_Uint64Value{Uint64Value: &types.UInt64Value{Value: 42}},
		},
		path: "uint64_value",
	},
	{
		name: "bool_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_BoolValue{BoolValue: &types.BoolValue{Value: true}},
		},
		path: "bool_value",
	},
	{
		name: "string_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_StringValue{StringValue: &types.StringValue{Value: "foo"}},
		},
		path: "string_value",
	},
	{
		name: "bytes_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_BytesValue{BytesValue: &types.BytesValue{Value: []byte("foo")}},
		},
		path: "bytes_value",
	},
	{
		name: "empty_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_EmptyValue{EmptyValue: &types.Empty{}},
		},
		path: "empty_value",
	},
	{
		name: "timestamp_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_TimestampValue{TimestampValue: mustTimestamp(testTime)},
		},
		path: "timestamp_value",
	},
	{
		name: "duration_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_DurationValue{DurationValue: mustDuration(testDuration)},
		},
		path: "duration_value",
	},
	{
		name: "field_mask_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_FieldMaskValue{FieldMaskValue: &types.FieldMask{Paths: []string{"foo.bar", "bar", "baz.qux"}}},
		},
		path: "field_mask_value",
	},
	{
		name: "value_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_ValueValue{ValueValue: &types.Value{Kind: &types.Value_StringValue{StringValue: "foo"}}},
		},
		path: "value_value",
	},
	{
		name: "list_value_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_ListValueValue{
				ListValueValue: &types.ListValue{
					Values: []*types.Value{
						{Kind: &types.Value_NullValue{}},
						{Kind: &types.Value_NumberValue{NumberValue: 12.34}},
						{Kind: &types.Value_StringValue{StringValue: "foo"}},
						{Kind: &types.Value_BoolValue{BoolValue: true}},
					},
				},
			},
		},
		path: "list_value_value",
	},
	{
		name: "struct_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_StructValue{
				StructValue: &types.Struct{
					Fields: map[string]*types.Value{
						"null":   {Kind: &types.Value_NullValue{}},
						"number": {Kind: &types.Value_NumberValue{NumberValue: 12.34}},
						"string": {Kind: &types.Value_StringValue{StringValue: "foo"}},
						"bool":   {Kind: &types.Value_BoolValue{BoolValue: true}},
					},
				},
			},
		},
		path: "struct_value",
	},
	{
		name: "any_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_AnyValue{
				AnyValue: mustAny(&SubMessage{Foo: "foo"}),
			},
		},
		path: "any_value",
	},
}

func TestMessageWithOneofWKTs(t *testing.T) {
	for _, tt := range testMessagesWithOneofWKTs {
		t.Run(tt.name, func(t *testing.T) {
			var dst MessageWithOneofWKTs

			err := dst.SetFields(tt.msg, tt.path)
			if err != nil {
				t.Fatal(err)
			}

			expectMessage(t, tt.msg, &dst)
		})
	}
}
