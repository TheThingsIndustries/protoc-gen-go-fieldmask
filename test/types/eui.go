// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
)

type EUI64 [8]byte

func (t EUI64) Marshal() ([]byte, error) {
	return t[:], nil
}

func (t *EUI64) MarshalTo(data []byte) (n int, err error) {
	return copy(data, t[:]), nil
}

func (t *EUI64) Unmarshal(data []byte) error {
	if len(data) != 8 {
		return fmt.Errorf("invalid data length: got %d, want 8", len(data))
	}
	var dto EUI64
	copy(dto[:], data)
	*t = dto
	return nil
}

func (t *EUI64) Size() int { return 8 }
