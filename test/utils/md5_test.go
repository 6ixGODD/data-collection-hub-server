package utils_test

import (
	"testing"

	"data-collection-hub-server/pkg/utils/crypt"
	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	md5Foo := crypt.MD5("foo")
	md5Bar := crypt.MD5("bar")

	assert.NotEmpty(t, md5Foo)
	assert.NotEmpty(t, md5Bar)
	assert.Equal(t, md5Foo, crypt.MD5("foo"))
	assert.Equal(t, md5Bar, crypt.MD5("bar"))
	assert.NotEqual(t, md5Foo, md5Bar)
}
