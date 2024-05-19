package crypt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tredoe/crypt"
	_ "github.com/tredoe/crypt/apr1_crypt"
)

func TestIsHashSupported(t *testing.T) {
	apr1 := crypt.IsHashSupported("$apr1$salt$hash")
	assert.True(t, apr1)
	other := crypt.IsHashSupported("$unknown$salt$hash")
	assert.False(t, other)
}
