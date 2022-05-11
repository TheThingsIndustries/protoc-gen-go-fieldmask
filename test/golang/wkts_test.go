// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package test_test

import (
	"testing"
	"time"

	. "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/test/golang"
	"google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	testTime     = time.Date(2006, time.January, 2, 15, 4, 5, 123456789, time.FixedZone("07:00", 7*3600))
	testDuration = time.Hour + 2*time.Minute + 3*time.Second + 123456789
)

func mustTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func mustDuration(t time.Duration) *durationpb.Duration {
	return durationpb.New(t)
}

func mustAny(pb proto.Message) *anypb.Any {
	any, err := anypb.New(pb)
	if err != nil {
		panic(err)
	}
	return any
}

var fullMessageWithWKTs = &MessageWithWKTs{
	DoubleValue: &wrapperspb.DoubleValue{Value: 12.34},
	DoubleValues: []*wrapperspb.DoubleValue{
		{Value: 12.34},
		{Value: 56.78},
	},
	FloatValue: &wrapperspb.FloatValue{Value: 12.34},
	FloatValues: []*wrapperspb.FloatValue{
		{Value: 12.34},
		{Value: 56.78},
	},
	Int32Value: &wrapperspb.Int32Value{Value: -42},
	Int32Values: []*wrapperspb.Int32Value{
		{Value: 1},
		{Value: 2},
		{Value: -42},
	},
	Int64Value: &wrapperspb.Int64Value{Value: -42},
	Int64Values: []*wrapperspb.Int64Value{
		{Value: 1},
		{Value: 2},
		{Value: -42},
	},
	Uint32Value: &wrapperspb.UInt32Value{Value: 42},
	Uint32Values: []*wrapperspb.UInt32Value{
		{Value: 1},
		{Value: 2},
		{Value: 42},
	},
	Uint64Value: &wrapperspb.UInt64Value{Value: 42},
	Uint64Values: []*wrapperspb.UInt64Value{
		{Value: 1},
		{Value: 2},
		{Value: 42},
	},
	BoolValue: &wrapperspb.BoolValue{Value: true},
	BoolValues: []*wrapperspb.BoolValue{
		{Value: true},
		{Value: false},
	},
	StringValue: &wrapperspb.StringValue{Value: "foo"},
	StringValues: []*wrapperspb.StringValue{
		{Value: "foo"},
		{Value: "bar"},
	},
	BytesValue: &wrapperspb.BytesValue{Value: []byte("foo")},
	BytesValues: []*wrapperspb.BytesValue{
		{Value: []byte("foo")},
		{Value: []byte("bar")},
	},
	EmptyValue:     &emptypb.Empty{},
	EmptyValues:    []*emptypb.Empty{{}, {}},
	TimestampValue: mustTimestamp(testTime),
	TimestampValues: []*timestamppb.Timestamp{
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
	DurationValues: []*durationpb.Duration{
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
	FieldMaskValue: &fieldmaskpb.FieldMask{Paths: []string{"foo.bar", "bar", "baz.qux"}},
	FieldMaskValues: []*fieldmaskpb.FieldMask{
		{Paths: []string{"foo.bar", "bar", "baz.qux"}},
	},
	ValueValue: &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: "foo"}},
	ValueValues: []*structpb.Value{
		{Kind: &structpb.Value_NullValue{}},
		{Kind: &structpb.Value_NumberValue{NumberValue: 12.34}},
		{Kind: &structpb.Value_StringValue{StringValue: "foo"}},
		{Kind: &structpb.Value_BoolValue{BoolValue: true}},
	},
	ListValueValue: &structpb.ListValue{
		Values: []*structpb.Value{
			{Kind: &structpb.Value_NullValue{}},
			{Kind: &structpb.Value_NumberValue{NumberValue: 12.34}},
			{Kind: &structpb.Value_StringValue{StringValue: "foo"}},
			{Kind: &structpb.Value_BoolValue{BoolValue: true}},
		},
	},
	ListValueValues: []*structpb.ListValue{
		{
			Values: []*structpb.Value{
				{Kind: &structpb.Value_NullValue{}},
				{Kind: &structpb.Value_NumberValue{NumberValue: 12.34}},
				{Kind: &structpb.Value_StringValue{StringValue: "foo"}},
				{Kind: &structpb.Value_BoolValue{BoolValue: true}},
			},
		},
	},
	StructValue: &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"null":   {Kind: &structpb.Value_NullValue{}},
			"number": {Kind: &structpb.Value_NumberValue{NumberValue: 12.34}},
			"string": {Kind: &structpb.Value_StringValue{StringValue: "foo"}},
			"bool":   {Kind: &structpb.Value_BoolValue{BoolValue: true}},
		},
	},
	StructValues: []*structpb.Struct{
		{Fields: map[string]*structpb.Value{"null": {Kind: &structpb.Value_NullValue{}}}},
		{Fields: map[string]*structpb.Value{"number": {Kind: &structpb.Value_NumberValue{NumberValue: 12.34}}}},
		{Fields: map[string]*structpb.Value{"string": {Kind: &structpb.Value_StringValue{StringValue: "foo"}}}},
		{Fields: map[string]*structpb.Value{"bool": {Kind: &structpb.Value_BoolValue{BoolValue: true}}}},
	},
	AnyValue: mustAny(&SubMessage{Foo: "foo"}),
	AnyValues: []*anypb.Any{
		mustAny(&SubMessage{Foo: "foo"}),
		mustAny(mustTimestamp(testTime)),
		mustAny(mustDuration(testDuration)),
		mustAny(&fieldmaskpb.FieldMask{Paths: []string{"foo.bar", "bar", "baz.qux"}}),
		mustAny(&structpb.Value{Kind: &structpb.Value_StringValue{StringValue: "foo"}}),
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
			Value: &MessageWithOneofWKTs_DoubleValue{DoubleValue: &wrapperspb.DoubleValue{Value: 12.34}},
		},
		path: "double_value",
	},
	{
		name: "float_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_FloatValue{FloatValue: &wrapperspb.FloatValue{Value: 12.34}},
		},
		path: "float_value",
	},
	{
		name: "int32_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_Int32Value{Int32Value: &wrapperspb.Int32Value{Value: -42}},
		},
		path: "int32_value",
	},
	{
		name: "int64_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_Int64Value{Int64Value: &wrapperspb.Int64Value{Value: -42}},
		},
		path: "int64_value",
	},
	{
		name: "uint32_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_Uint32Value{Uint32Value: &wrapperspb.UInt32Value{Value: 42}},
		},
		path: "uint32_value",
	},
	{
		name: "uint64_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_Uint64Value{Uint64Value: &wrapperspb.UInt64Value{Value: 42}},
		},
		path: "uint64_value",
	},
	{
		name: "bool_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_BoolValue{BoolValue: &wrapperspb.BoolValue{Value: true}},
		},
		path: "bool_value",
	},
	{
		name: "string_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_StringValue{StringValue: &wrapperspb.StringValue{Value: "foo"}},
		},
		path: "string_value",
	},
	{
		name: "bytes_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_BytesValue{BytesValue: &wrapperspb.BytesValue{Value: []byte("foo")}},
		},
		path: "bytes_value",
	},
	{
		name: "empty_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_EmptyValue{EmptyValue: &emptypb.Empty{}},
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
			Value: &MessageWithOneofWKTs_FieldMaskValue{FieldMaskValue: &fieldmaskpb.FieldMask{Paths: []string{"foo.bar", "bar", "baz.qux"}}},
		},
		path: "field_mask_value",
	},
	{
		name: "value_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_ValueValue{ValueValue: &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: "foo"}}},
		},
		path: "value_value",
	},
	{
		name: "list_value_value",
		msg: &MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_ListValueValue{
				ListValueValue: &structpb.ListValue{
					Values: []*structpb.Value{
						{Kind: &structpb.Value_NullValue{}},
						{Kind: &structpb.Value_NumberValue{NumberValue: 12.34}},
						{Kind: &structpb.Value_StringValue{StringValue: "foo"}},
						{Kind: &structpb.Value_BoolValue{BoolValue: true}},
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
				StructValue: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"null":   {Kind: &structpb.Value_NullValue{}},
						"number": {Kind: &structpb.Value_NumberValue{NumberValue: 12.34}},
						"string": {Kind: &structpb.Value_StringValue{StringValue: "foo"}},
						"bool":   {Kind: &structpb.Value_BoolValue{BoolValue: true}},
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
