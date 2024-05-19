// SPDX-FileCopyrightText: 2013 Jonás Melián
// SPDX-License-Identifier: BSD-2-Clause

package crypt

import "errors"

var (
	ErrKeyMismatch = errors.New("given password does not match")
	ErrUnknown     = errors.New("unknown crypt function")
)
