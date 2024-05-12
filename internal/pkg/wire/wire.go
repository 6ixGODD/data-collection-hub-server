//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	"data-collection-hub-server/internal/app"
	"data-collection-hub-server/internal/pkg/api/v1"
	adminapi "data-collection-hub-server/internal/pkg/api/v1/admin"
	commonapi "data-collection-hub-server/internal/pkg/api/v1/common"
	userapi "data-collection-hub-server/internal/pkg/api/v1/user"
	"data-collection-hub-server/internal/pkg/dal"
	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/internal/pkg/service/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	"data-collection-hub-server/internal/pkg/service/common"
	commonservice "data-collection-hub-server/internal/pkg/service/common/mods"
	"data-collection-hub-server/internal/pkg/service/user"
	userservice "data-collection-hub-server/internal/pkg/service/user/mods"
	"data-collection-hub-server/pkg/jwt"
	"data-collection-hub-server/pkg/middleware"
	ware "data-collection-hub-server/pkg/middleware/mods"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
	"github.com/google/wire"
)

var DalSet = wire.NewSet(
	mongo.New,
	redis.New,
	zap.New,
	dal.New,
	dao.NewDocumentationDao,
	dao.NewUserDao,
	dao.NewNoticeDao,
	dao.NewInstructionDataDao,
	dao.NewOperationLogDao,
	dao.NewLoginLogDao,
	dao.NewErrorLogDao,
)

var ServiceSet = wire.NewSet(
	service.New,
	admin.New,
	adminservice.NewDataAuditService,
	adminservice.NewDocumentationService,
	adminservice.NewLogsService,
	adminservice.NewNoticeService,
	adminservice.NewUserService,
	adminservice.NewStatisticService,
	common.New,
	commonservice.NewNoticeService,
	commonservice.NewDocumentationService,
	commonservice.NewAuthService,
	user.New,
	userservice.NewDatasetService,
	userservice.NewStatisticService,
)

var MiddlewareSet = wire.NewSet(
	middleware.New,
	zap.New,
	jwt.New,
	ware.NewAuthMiddleware,
	ware.NewLoggingMiddleware,
	ware.NewPrometheusMiddleware,
)

var ApiSet = wire.NewSet(
	api.New,
	adminapi.New,
	userapi.New,
	commonapi.New,
)

func InitializeApp(ctx context.Context) *app.App {
	panic(
		wire.Build(
			mongo.New,
			redis.New,
			zap.New,
			DalSet,
			ServiceSet,
			MiddlewareSet,
			ApiSet,
			app.New,
		),
	)
}
