// SPDX-FileCopyrightText: 2013 Jonás Melián
// SPDX-License-Identifier: BSD-2-Clause

package crypt

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrKeyMismatch = errors.New("given password does not match")
	ErrUnknown     = errors.New("unknown crypt function")
)

// UnknownError reports a hashed string without a known crypt function.
type UnknownError string

func (e UnknownError) Error() string {
	// Hash like: $1$deadbeef$...
	fields := strings.Split(string(e), "$")

	return fmt.Sprintf("%s: $%s$", ErrUnknown, fields[1])
}
