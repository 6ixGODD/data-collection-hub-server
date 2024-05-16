package utils__test

import (
	"testing"

	"data-collection-hub-server/pkg/utils/crypt"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hash, err := crypt.PasswordHash("foo")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	assert.True(t, crypt.VerifyPassword("foo", hash))
	assert.False(t, crypt.VerifyPassword("bar", hash))
}
