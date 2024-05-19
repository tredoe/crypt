// SPDX-FileCopyrightText: 2012 Jeramey Crawford <jeramey@antihe.ro>
// SPDX-License-Identifier: BSD-2-Clause

// Package apr1_crypt implements the standard Unix MD5-crypt algorithm created
// by Poul-Henning Kamp for FreeBSD, and modified by the Apache project.
//
// The only change from MD5-crypt is the use of the magic constant "$apr1$"
// instead of "$1$". The algorithms are otherwise identical.
package apr1_crypt

import (
	"github.com/tredoe/crypt"
	"github.com/tredoe/crypt/common"
	"github.com/tredoe/crypt/md5_crypt"
)

func init() {
	crypt.RegisterCrypt(crypt.APR1, New, MagicPrefix)
}

const (
	MagicPrefix   = "$apr1$"
	SaltLenMin    = 1
	SaltLenMax    = 8
	RoundsDefault = 1000
)

// New returns a new crypt.Crypter computing the variant "apr1" of MD5-crypt
func New() crypt.Crypter {
	crypter := md5_crypt.New()
	crypter.SetSalt(common.Salt{
		MagicPrefix:   []byte(MagicPrefix),
		SaltLenMin:    SaltLenMin,
		SaltLenMax:    SaltLenMax,
		RoundsDefault: RoundsDefault,
	})
	return crypter
}
