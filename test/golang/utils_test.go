// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package test_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	proto "google.golang.org/protobuf/proto"
)

func expectMessage(t *testing.T, expected, got proto.Message) {
	t.Helper()

	expectedMsgText := prototext.Format(expected)
	gotMsgText := prototext.Format(got)

	generatedDiff := cmp.Diff(expectedMsgText, gotMsgText)

	if generatedDiff != "" {
		t.Errorf("expected : %s", expectedMsgText)
		t.Errorf("got: %s", gotMsgText)
		if generatedDiff != "" {
			t.Errorf("  diff   : %s", generatedDiff)
		}
	}
}
