package utils__test

import (
	"testing"

	"data-collection-hub-server/pkg/utils/crypt"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	a := assert.New(t)

	// Test1 hash
	hash, err := crypt.PasswordHash("123456")
	a.NoError(err)
	a.NotEmpty(hash)

	// Test1 verify
	a.True(crypt.VerifyPassword("123456", hash))
	a.False(crypt.VerifyPassword("1234567", hash))
}
