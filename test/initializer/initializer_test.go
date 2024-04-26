package initializer

import (
	"testing"

	"data-collection-hub-server/core/config"
	"github.com/stretchr/testify/assert"
)

func TestInitializer(t *testing.T) {
	a := assert.New(t)
	// ctx := context.Background()
	// app, err := initializer.NewApp(ctx)
	// if err != nil {
	// 	return
	// }
	// // Environment variables
	// a.NotNil(os.Getenv("APP_PORT"))
	// a.Equal("3000", os.Getenv("APP_PORT"))
	// a.NotNil(os.Getenv("APP_NAME"))
	// a.NotNil(os.Getenv("APP_VERSION"))
	// a.NotNil(os.Getenv("APP_HOST"))
	// err = app.Listen(":3000")
	// if err != nil {
	// 	return
	// }
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	} else {
		a.NotNil(cfg)
		t.Logf("config: %+v", cfg)
	}

}
