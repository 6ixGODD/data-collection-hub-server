package wire

import (
	"context"

	"data-collection-hub-server/internal/app"
	"data-collection-hub-server/internal/pkg/api/v1"
	"data-collection-hub-server/internal/pkg/api/v1/admin"
	"data-collection-hub-server/internal/pkg/api/v1/common"
	"data-collection-hub-server/internal/pkg/api/v1/user"
	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/router"
	v1 "data-collection-hub-server/internal/pkg/router/v1"
	"github.com/google/wire"
)

func InitializeApp(ctx context.Context) (*app.App, error) {
	wire.Build(
		config.New,
		InitializeMongo,
		InitializeRedis,
		InitializeZap,
		InitializeJwt,
		InitializePrometheus,
		wire.NewSet(wire.Struct(new(app.App), "*")),
		wire.NewSet(wire.Struct(new(router.Router), "*")),
		wire.NewSet(wire.Struct(new(v1.Router), "*")),
		api.New,
		admin.New,
		common.New,
		user.New,
	)
	return new(app.App), nil
}
