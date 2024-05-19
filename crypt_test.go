// SPDX-FileCopyrightText: 2013 Jonás Melián
// SPDX-License-Identifier: BSD-2-Clause

package crypt_test

import (
	"testing"

	"github.com/tredoe/crypt"
	_ "github.com/tredoe/crypt/apr1_crypt"
)

func TestSupport(t *testing.T) {
	hash := "$apr1$salt$hash"
	ok := crypt.IsHashSupported(hash)
	if !ok {
		t.Errorf("expect support for hash: %q", hash)
	}

	hash = "$unknown$salt$hash"
	ok = crypt.IsHashSupported(hash)
	if ok {
		t.Errorf("expect no support for hash: %q", hash)
	}
}
