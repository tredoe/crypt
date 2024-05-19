// SPDX-FileCopyrightText: 2013 Jonás Melián
// SPDX-License-Identifier: BSD-2-Clause

package crypt_test

import (
	"strings"
	"testing"

	"github.com/tredoe/crypt"
	_ "github.com/tredoe/crypt/apr1_crypt"
)

func TestSupport(t *testing.T) {
	hash := "$apr1$salt$hash"
	_, err := crypt.NewFromHash(hash)
	if err != nil {
		t.Errorf("expect support for hash: %q", hash)
	}

	hash = "$unknown$salt$hash"
	if _, err = crypt.NewFromHash(hash); err == nil {
		t.Fatalf("expect no support for hash: %q", hash)
	}
	if !strings.HasSuffix(err.Error(), "$unknown$") {
		t.Error("expect that error got the crypt magic identifier")
	}
}
