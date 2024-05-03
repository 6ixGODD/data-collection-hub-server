package wire

import (
	"data-collection-hub-server/internal/app"
	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/internal/pkg/service/admin_service/modules"
	modules2 "data-collection-hub-server/internal/pkg/service/common_service/modules"
	modules3 "data-collection-hub-server/internal/pkg/service/user_service/modules"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
	"github.com/google/wire"
)

func InitializeApp() (*app.App, error) {
	wire.Build(app.New, zap.New, redis.New, mongo.New, config.New)
	return &app.App{}, nil
}

var appSet = wire.NewSet(
	app.New,
)

var configSet = wire.NewSet(
	config.New,
)

var mongoSet = wire.NewSet(
	mongo.New,
)

var redisSet = wire.NewSet(
	redis.New,
)

var zapSet = wire.NewSet(
	zap.New,
)

var serviceSet = wire.NewSet(
	service.NewCore,
	modules.NewDataAuditService,
	wire.Bind(new(modules.DataAuditService), new(*modules.DataAuditServiceImpl)),
	modules.NewDocumentationService,
	wire.Bind(new(modules.DocumentationService), new(*modules.DocumentationServiceImpl)),
	modules.NewLogsService,
	wire.Bind(new(modules.LogsService), new(*modules.LogsServiceImpl)),
	modules.NewNoticeService,
	wire.Bind(new(modules.NoticeService), new(*modules.NoticeServiceImpl)),
	modules.NewStatisticService,
	wire.Bind(new(modules.StatisticService), new(*modules.StatisticServiceImpl)),
	modules.NewUserService,
	wire.Bind(new(modules.UserService), new(*modules.UserServiceImpl)),
	modules3.NewDatasetService,
	wire.Bind(new(modules3.DatasetService), new(*modules3.DatasetServiceImpl)),
	modules3.NewStatisticService,
	wire.Bind(new(modules3.StatisticService), new(*modules3.StatisticServiceImpl)),
	modules2.NewAuthService,
	wire.Bind(new(modules2.AuthService), new(*modules2.AuthServiceImpl)),
	modules2.NewDocumentationService,
	wire.Bind(new(modules2.DocumentationService), new(*modules2.DocumentationServiceImpl)),
	modules2.NewNoticeService,
	wire.Bind(new(modules2.NoticeService), new(*modules2.NoticeServiceImpl)),
	modules2.NewProfileService,
	wire.Bind(new(modules2.ProfileService), new(*modules2.ProfileServiceImpl)),
)
