package initializer

import (
	"os"
	"testing"

	"data-collection-hub-server/initializer"
	"github.com/stretchr/testify/assert"
)

func TestInitializer(t *testing.T) {
	a := assert.New(t)
	initializer.InitApp()
	// Environment variables
	a.NotNil(os.Getenv("APP_PORT"))
	a.Equal("3000", os.Getenv("APP_PORT"))
	a.NotNil(os.Getenv("APP_NAME"))
	a.NotNil(os.Getenv("APP_VERSION"))
	a.NotNil(os.Getenv("APP_HOST"))
}
