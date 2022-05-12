// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package test_test

import (
	"testing"

	. "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/test/gogo"
)

var fullMessageWithSubMessages = &MessageWithSubMessages{
	A: &SubMessage{
		Foo: "foo",
		Bar: "bar",
	},
	B: &SubMessage{
		Foo: "foo",
		Bar: "bar",
	},
}

func TestMessageWithSubMessages(t *testing.T) {
	var dst MessageWithSubMessages

	expectStringSlice(t, nil, dst.FieldPaths(0))
	expectStringSlice(t, []string{"a", "b"}, dst.FieldPaths(1))
	expectStringSlice(t, []string{"a", "a.foo", "a.bar", "b", "b.foo", "b.bar"}, dst.FieldPaths(2))
	expectStringSlice(t, []string{"a", "a.foo", "a.bar", "b", "b.foo", "b.bar"}, dst.FieldPaths(3))

	normalized, err := dst.NormalizeFieldPaths("a", "a.foo", "a.bar", "b", "b.foo", "b.bar")
	if err != nil {
		t.Errorf("unexpected error in NormalizeFieldPaths: %v", err)
	}
	expectStringSlice(t, []string{"a", "b"}, normalized)

	err = dst.SetFields(
		fullMessageWithSubMessages,
		"a",
		"b.foo",
	)
	if err != nil {
		t.Fatal(err)
	}

	expectMessage(t, &MessageWithSubMessages{
		A: &SubMessage{
			Foo: "foo",
			Bar: "bar",
		},
		B: &SubMessage{
			Foo: "foo",
		},
	}, &dst)
}

var testMessagesWithOneofSubMessages = []struct {
	name     string
	src      *MessageWithOneofSubMessages
	dst      *MessageWithOneofSubMessages
	expected *MessageWithOneofSubMessages
	path     string
}{
	{
		name: "empty",
		src:  &MessageWithOneofSubMessages{},
		path: "sub",
	},
	{
		name: "a",
		src: &MessageWithOneofSubMessages{
			Sub: &MessageWithOneofSubMessages_A{A: &SubMessage{Foo: "foo"}},
		},
		path: "a",
	},
	{
		name: "b",
		src: &MessageWithOneofSubMessages{
			Sub: &MessageWithOneofSubMessages_B{B: &SubMessage{Foo: "foo"}},
		},
		path: "b",
	},
	{
		name: "a.bar",
		src: &MessageWithOneofSubMessages{
			Sub: &MessageWithOneofSubMessages_A{A: &SubMessage{Foo: "bar", Bar: "qux"}},
		},
		dst: &MessageWithOneofSubMessages{
			Sub: &MessageWithOneofSubMessages_A{A: &SubMessage{Foo: "foo"}},
		},
		expected: &MessageWithOneofSubMessages{
			Sub: &MessageWithOneofSubMessages_A{A: &SubMessage{Foo: "foo", Bar: "qux"}},
		},
		path: "a.bar",
	},
}

func TestMessageWithOneofSubMessages(t *testing.T) {
	var dst MessageWithOneofSubMessages

	expectStringSlice(t, nil, dst.FieldPaths(0))
	expectStringSlice(t, []string{"a", "b"}, dst.FieldPaths(1))
	expectStringSlice(t, []string{"a", "a.foo", "a.bar", "b", "b.foo", "b.bar"}, dst.FieldPaths(2))
	expectStringSlice(t, []string{"a", "a.foo", "a.bar", "b", "b.foo", "b.bar"}, dst.FieldPaths(3))

	normalized, err := dst.NormalizeFieldPaths("a", "a.foo", "a.bar", "b", "b.foo", "b.bar")
	if err != nil {
		t.Errorf("unexpected error in NormalizeFieldPaths: %v", err)
	}
	expectStringSlice(t, []string{"a", "b"}, normalized)

	for _, tt := range testMessagesWithOneofSubMessages {
		t.Run(tt.name, func(t *testing.T) {
			dst := tt.dst
			if dst == nil {
				dst = &MessageWithOneofSubMessages{}
			}

			err := dst.SetFields(tt.src, tt.path)
			if err != nil {
				t.Fatal(err)
			}

			expected := tt.expected
			if expected == nil {
				expected = tt.src
			}

			expectMessage(t, expected, dst)
		})
	}
}
