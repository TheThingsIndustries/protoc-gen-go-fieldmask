// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package test_test

import (
	"testing"

	. "github.com/TheThingsIndustries/protoc-gen-go-field-setters/test/golang"
)

var fullMessageWithGoGoOptions = &MessageWithGoGoOptions{
	EuiWithCustomName:                   []byte{1, 2, 3, 4, 5, 6, 7, 8},
	EuiWithCustomNameAndType:            []byte{1, 2, 3, 4, 5, 6, 7, 8},
	NonNullableEuiWithCustomNameAndType: []byte{1, 2, 3, 4, 5, 6, 7, 8},
	EuisWithCustomNameAndType: [][]byte{
		{1, 2, 3, 4, 5, 6, 7, 8},
	},
	Duration:             mustDuration(testDuration),
	NonNullableDuration:  mustDuration(testDuration),
	Timestamp:            mustTimestamp(testTime),
	NonNullableTimestamp: mustTimestamp(testTime),
}

func TestMessageWithGoGoOptions(t *testing.T) {
	var dst MessageWithGoGoOptions

	err := dst.SetFields(
		fullMessageWithGoGoOptions,
		"eui_with_custom_name",
		"eui_with_custom_name_and_type",
		"non_nullable_eui_with_custom_name_and_type",
		"euis_with_custom_name_and_type",
		"duration",
		"non_nullable_duration",
		"timestamp",
		"non_nullable_timestamp",
	)
	if err != nil {
		t.Fatal(err)
	}

	expectMessage(t, fullMessageWithGoGoOptions, &dst)
}

var fullMessageWithNullable = &MessageWithNullable{
	Sub: &SubMessage{Foo: "foo"},
	Subs: []*SubMessage{
		{Foo: "foo"},
	},
}

func TestMessageWithNullable(t *testing.T) {
	for _, paths := range [][]string{
		{"sub", "subs"},
		{"sub.foo", "subs"},
	} {
		var dst MessageWithNullable

		err := dst.SetFields(
			fullMessageWithNullable,
			paths...,
		)
		if err != nil {
			t.Fatal(err)
		}

		expectMessage(t, fullMessageWithNullable, &dst)
	}
}

var fullMessageWithEmbedded = &MessageWithEmbedded{
	Sub: &SubMessage{Foo: "foo"},
}

func TestMessageWithEmbedded(t *testing.T) {
	for _, paths := range [][]string{
		{"sub"},
		{"sub.foo"},
	} {
		var dst MessageWithEmbedded

		err := dst.SetFields(
			fullMessageWithEmbedded,
			paths...,
		)
		if err != nil {
			t.Fatal(err)
		}

		expectMessage(t, fullMessageWithEmbedded, &dst)
	}
}

var fullMessageWithNullableEmbedded = &MessageWithNullableEmbedded{
	Sub: &SubMessage{Foo: "foo"},
}

func TestMessageWithNullableEmbedded(t *testing.T) {
	for _, paths := range [][]string{
		{"sub"},
		{"sub.foo"},
	} {
		var dst MessageWithNullableEmbedded

		err := dst.SetFields(
			fullMessageWithNullableEmbedded,
			paths...,
		)
		if err != nil {
			t.Fatal(err)
		}

		expectMessage(t, fullMessageWithNullableEmbedded, &dst)
	}
}
