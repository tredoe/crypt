// SPDX-FileCopyrightText: 2013 Jonás Melián
// SPDX-License-Identifier: BSD-2-Clause

// Package crypt provides interface for password crypt functions and collects
// common constants.
package crypt

import (
	"strings"

	"github.com/tredoe/crypt/common"
)

// Crypter is the common interface implemented by all crypt functions.
type Crypter interface {
	// Generate performs the hashing algorithm, returning a full hash suitable
	// for storage and later password verification.
	//
	// If the salt is empty, a randomly-generated salt will be generated with a
	// length of SaltLenMax and number RoundsDefault of rounds.
	//
	// Any error only can be got when the salt argument is not empty.
	Generate(key, salt []byte) (string, error)

	// Verify compares a hashed key with its possible key equivalent.
	// Returns nil on success, or an error on failure; if the hashed key is
	// diffrent, the error is "ErrKeyMismatch".
	Verify(hashedKey string, key []byte) error

	// Cost returns the hashing cost (in rounds) used to create the given hashed
	// key.
	//
	// When, in the future, the hashing cost of a key needs to be increased in
	// order to adjust for greater computational power, this function allows one
	// to establish which keys need to be updated.
	//
	// The algorithms based in MD5-crypt use a fixed value of rounds.
	Cost(hashedKey string) (int, error)

	// SetSalt sets a different salt. It is used to easily create derivated
	// algorithms, i.e. "apr1_crypt" from "md5_crypt".
	SetSalt(salt common.Salt)
}

// Crypt identifies a crypt function that is implemented in another package.
type Crypt uint

const (
	APR1   Crypt = iota // import github.com/tredoe/crypt/apr1_crypt
	MD5                 // import github.com/tredoe/crypt/md5_crypt
	SHA256              // import github.com/tredoe/crypt/sha256_crypt
	SHA512              // import github.com/tredoe/crypt/sha512_crypt
	maxCrypt
)

var (
	crypts        = make([]func() Crypter, maxCrypt)
	cryptPrefixes = make([]string, maxCrypt)
)

// * * *

// New returns a new crypter.
func New(c Crypt) Crypter { return c.New() }

// NewFromHash returns a new Crypter using the prefix in the given hashed key.
func NewFromHash(hashedKey string) (Crypter, error) {
	for i := range cryptPrefixes {
		prefix := cryptPrefixes[i]

		if crypts[i] != nil && strings.HasPrefix(hashedKey, prefix) {
			c := Crypt(uint(i))
			return c.New(), nil
		}
	}

	return nil, ErrUnknown
}

// New returns new Crypter making the Crypt c.
// New panics if the Crypt c is unavailable.
func (c Crypt) New() Crypter {
	if c < maxCrypt {
		f := crypts[c]
		if f != nil {
			return f()
		}
	}
	panic(ErrUnknown)
}

// Available reports whether the Crypt c is available.
func (c Crypt) Available() bool {
	return c < maxCrypt && crypts[c] != nil
}

// RegisterCrypt registers a function that returns a new instance of the given
// crypt function. This is intended to be called from the init function in
// packages that implement crypt functions.
func RegisterCrypt(c Crypt, f func() Crypter, prefix string) {
	if c >= maxCrypt {
		panic(ErrUnknown)
	}
	crypts[c] = f
	cryptPrefixes[c] = prefix
}
