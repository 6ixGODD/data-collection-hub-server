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
	"data-collection-hub-server/internal/pkg/dal"
	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/router"
	routerv1 "data-collection-hub-server/internal/pkg/router/v1"
	routers "data-collection-hub-server/internal/pkg/router/v1/mods"
	"data-collection-hub-server/internal/pkg/service"
	adminservice "data-collection-hub-server/internal/pkg/service/admin"
	adminservices "data-collection-hub-server/internal/pkg/service/admin/mods"
	commonservice "data-collection-hub-server/internal/pkg/service/common"
	commonservices "data-collection-hub-server/internal/pkg/service/common/mods"
	userservice "data-collection-hub-server/internal/pkg/service/user"
	userservices "data-collection-hub-server/internal/pkg/service/user/mods"
	"data-collection-hub-server/pkg/cron"
	"data-collection-hub-server/pkg/middleware"
	middlewares "data-collection-hub-server/pkg/middleware/mods"
	"github.com/google/wire"
)

var (
	RouterSet = wire.NewSet(
		wire.Struct(new(routers.AdminRouter), "*"),
		wire.Struct(new(routers.UserRouter), "*"),
		wire.Struct(new(routers.CommonRouter), "*"),
		wire.Struct(new(routerv1.Router), "*"),
		wire.Struct(new(router.Router), "*"),
	)

	ApiSet = wire.NewSet(
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

	ServiceSet = wire.NewSet(
		wire.Struct(new(service.Service), "*"),
		wire.Struct(new(adminservice.Admin), "*"),
		wire.Struct(new(userservice.User), "*"),
		wire.Struct(new(commonservice.Common), "*"),
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
	)

	DaoSet = wire.NewSet(
		wire.Struct(new(dal.Dao), "*"),
		dao.NewUserDao,
		dao.NewInstructionDataDao,
		dao.NewNoticeDao,
		dao.NewLoginLogDao,
		dao.NewOperationLogDao,
		dao.NewErrorLogDao,
		dao.NewDocumentationDao,
	)

	MiddlewareSet = wire.NewSet(
		middlewares.NewLoggingMiddleware,
		middlewares.NewPrometheusMiddleware,
		middlewares.NewAuthMiddleware,
		wire.Struct(new(middleware.Middleware), "*"),
	)

	SchedulerSet = wire.NewSet(
		cron.New,
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
		DaoSet,
		ServiceSet,
		ApiSet,
		RouterSet,
		MiddlewareSet,
		SchedulerSet,
		app.New,
	)
	return new(app.App), nil
}
