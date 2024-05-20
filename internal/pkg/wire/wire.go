//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	"data-collection-hub-server/internal/app"
	"data-collection-hub-server/internal/pkg/api/v1"
	adminapi "data-collection-hub-server/internal/pkg/api/v1/admin"
	adminapis "data-collection-hub-server/internal/pkg/api/v1/admin/mods"
	commonapi "data-collection-hub-server/internal/pkg/api/v1/common"
	commonapis "data-collection-hub-server/internal/pkg/api/v1/common/mods"
	userapi "data-collection-hub-server/internal/pkg/api/v1/user"
	userapis "data-collection-hub-server/internal/pkg/api/v1/user/mods"
	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	daos "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/middleware"
	wares "data-collection-hub-server/internal/pkg/middleware/mods"
	"data-collection-hub-server/internal/pkg/router"
	routerv1 "data-collection-hub-server/internal/pkg/router/v1"
	routers "data-collection-hub-server/internal/pkg/router/v1/mods"
	"data-collection-hub-server/internal/pkg/service"
	adminservice "data-collection-hub-server/internal/pkg/service/admin"
	adminservices "data-collection-hub-server/internal/pkg/service/admin/mods"
	commonservice "data-collection-hub-server/internal/pkg/service/common"
	commonservices "data-collection-hub-server/internal/pkg/service/common/mods"
	sysservice "data-collection-hub-server/internal/pkg/service/sys"
	sysservices "data-collection-hub-server/internal/pkg/service/sys/mods"
	userservice "data-collection-hub-server/internal/pkg/service/user"
	userservices "data-collection-hub-server/internal/pkg/service/user/mods"
	"data-collection-hub-server/internal/pkg/tasks"
	"github.com/google/wire"
)

var (
	RouterProviderSet = wire.NewSet(
		wire.Struct(new(routers.AdminRouter), "*"),
		wire.Struct(new(routers.UserRouter), "*"),
		wire.Struct(new(routers.CommonRouter), "*"),
		wire.Struct(new(routerv1.Router), "*"),
		wire.Struct(new(router.Router), "*"),
	)

	ApiProviderSet = wire.NewSet(
		wire.Struct(new(commonapis.AuthApi), "*"),
		wire.Struct(new(commonapis.ProfileApi), "*"),
		wire.Struct(new(commonapis.DocumentationApi), "*"),
		wire.Struct(new(commonapis.NoticeApi), "*"),
		wire.Struct(new(userapis.DatasetApi), "*"),
		wire.Struct(new(userapis.StatisticApi), "*"),
		wire.Struct(new(adminapis.UserApi), "*"),
		wire.Struct(new(adminapis.DocumentationApi), "*"),
		wire.Struct(new(adminapis.NoticeApi), "*"),
		wire.Struct(new(adminapis.StatisticApi), "*"),
		wire.Struct(new(adminapis.LogsApi), "*"),
		wire.Struct(new(adminapis.DataAuditApi), "*"),
		wire.Struct(new(commonapi.Common), "*"),
		wire.Struct(new(userapi.User), "*"),
		wire.Struct(new(adminapi.Admin), "*"),
		wire.Struct(new(api.Api), "*"),
	)

	ServiceProviderSet = wire.NewSet(
		wire.Struct(new(service.Core), "*"),
		wire.Struct(new(adminservice.Admin), "*"),
		wire.Struct(new(userservice.User), "*"),
		wire.Struct(new(commonservice.Common), "*"),
		wire.Struct(new(sysservice.Sys), "*"),
		adminservices.NewDataAuditService,
		adminservices.NewStatisticService,
		adminservices.NewUserService,
		adminservices.NewNoticeService,
		adminservices.NewDocumentationService,
		adminservices.NewLogsService,
		commonservices.NewAuthService,
		commonservices.NewProfileService,
		commonservices.NewDocumentationService,
		commonservices.NewNoticeService,
		userservices.NewDatasetService,
		userservices.NewStatisticService,
		sysservices.NewLogsService,
	)

	DaoProviderSet = wire.NewSet(
		dao.NewCore,
		dao.NewCache,
		daos.NewUserDao,
		daos.NewInstructionDataDao,
		daos.NewNoticeDao,
		daos.NewLoginLogDao,
		daos.NewOperationLogDao,
		daos.NewDocumentationDao,
	)

	MiddlewareProviderSet = wire.NewSet(
		wire.Struct(new(wares.LoggingMiddleware), "*"),
		wire.Struct(new(wares.PrometheusMiddleware), "*"),
		wire.Struct(new(wares.AuthMiddleware), "*"),
		wire.Struct(new(middleware.Middleware), "*"),
	)

	SchedulerProviderSet = wire.NewSet(
		tasks.New,
	)
)

// InitializeApp initialize app
func InitializeApp(ctx context.Context) (*app.App, error) {
	wire.Build(
		config.New,
		InitializeMongo,
		InitializeRedis,
		InitializeZap,
		InitializeJwt,
		InitializePrometheus,
		DaoProviderSet,
		ServiceProviderSet,
		ApiProviderSet,
		RouterProviderSet,
		MiddlewareProviderSet,
		SchedulerProviderSet,
		app.New,
	)
	return new(app.App), nil
}
