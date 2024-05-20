package utils__test

import (
	"testing"

	"data-collection-hub-server/pkg/utils/crypt"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hash, err := crypt.Hash("foo")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	assert.True(t, crypt.VerifyHash("foo", hash))
	assert.False(t, crypt.VerifyHash("bar", hash))
}
