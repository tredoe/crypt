// SPDX-FileCopyrightText: 2013 Jonás Melián
// SPDX-License-Identifier: BSD-2-Clause

package crypt

import (
	"errors"
	"fmt"
)

var (
	ErrKeyMismatch = errors.New("given password does not match")
	ErrUnknown     = errors.New("unknown crypt function")
)

// UnknownError reports a string without a known crypt function.
type UnknownError string

func (e UnknownError) Error() string {
	return fmt.Sprintf("%s: %s", ErrUnknown, string(e))
}
