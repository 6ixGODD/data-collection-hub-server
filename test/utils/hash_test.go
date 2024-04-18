package utils_test

import (
	"testing"

	"data-collection-hub-server/utils/crypt"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	a := assert.New(t)

	// Test hash
	hash, _ := crypt.PasswordHash("123456")
	a.NotEmpty(hash)

	// Test verify
	a.True(crypt.VerifyPassword("123456", hash))
	a.False(crypt.VerifyPassword("1234567", hash))
}
