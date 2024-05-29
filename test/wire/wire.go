//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	daos "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/service"
	adminservice "data-collection-hub-server/internal/pkg/service/admin"
	adminservices "data-collection-hub-server/internal/pkg/service/admin/mods"
	commonservice "data-collection-hub-server/internal/pkg/service/common"
	commonservices "data-collection-hub-server/internal/pkg/service/common/mods"
	sysservice "data-collection-hub-server/internal/pkg/service/sys"
	sysservices "data-collection-hub-server/internal/pkg/service/sys/mods"
	userservice "data-collection-hub-server/internal/pkg/service/user"
	userservices "data-collection-hub-server/internal/pkg/service/user/mods"
	"data-collection-hub-server/pkg/jwt"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/prometheus"
	"data-collection-hub-server/pkg/redis"
	logging "data-collection-hub-server/pkg/zap"
	"data-collection-hub-server/test/mock"
	"github.com/google/wire"
)

type Injector struct {
	// Common
	Ctx        context.Context
	Config     *config.Config
	Cache      *dao.Cache
	Mongo      *mongo.Mongo
	Redis      *redis.Redis
	Zap        *logging.Zap
	Jwt        *jwt.Jwt
	Prometheus *prometheus.Prometheus

	// DAOs
	UserDao            daos.UserDao
	InstructionDataDao daos.InstructionDataDao
	NoticeDao          daos.NoticeDao
	DocumentationDao   daos.DocumentationDao
	LoginLogDao        daos.LoginLogDao
	OperationLogDao    daos.OperationLogDao

	// Mocks for DAOs
	UserDaoMock            *mock.UserDaoMock
	InstructionDataDaoMock *mock.InstructionDataDaoMock
	NoticeDaoMock          *mock.NoticeDaoMock
	DocumentationDaoMock   *mock.DocumentationDaoMock
	LoginLogDaoMock        *mock.LoginLogDaoMock
	OperationLogDaoMock    *mock.OperationLogDaoMock

	// Services
	// Admin services
	AdminDataAuditService     adminservices.DataAuditService
	AdminDocumentationService adminservices.DocumentationService
	AdminNoticeService        adminservices.NoticeService
	AdminLogsService          adminservices.LogsService
	AdminStatisticService     adminservices.StatisticService
	AdminUserService          adminservices.UserService
	// Common services
	CommonAuthService          commonservices.AuthService
	CommonIdempotencyService   commonservices.IdempotencyService
	CommonDocumentationService commonservices.DocumentationService
	CommonNoticeService        commonservices.NoticeService
	CommonProfileService       commonservices.ProfileService
	// Sys services
	SysLogsService sysservices.LogsService
	// User services
	UserDatasetService   userservices.DatasetService
	UserStatisticService userservices.StatisticService
}

var (
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
		commonservices.NewIdempotencyService,
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

	MockProviderSet = wire.NewSet(
		mock.NewUserDaoMockWithRandomData,
		mock.NewInstructionDataDaoMockWithRandomData,
		mock.NewNoticeDaoMockWithRandomData,
		mock.NewLoginLogDaoMockWithRandomData,
		mock.NewOperationLogDaoMockWithRandomData,
		mock.NewDocumentationDaoMockWithRandomData,
	)
)

func InitializeTestInjector(ctx context.Context, config *config.Config, n int) (*Injector, error) {
	wire.Build(
		InitializeMongo,
		InitializeRedis,
		InitializeZap,
		InitializeJwt,
		InitializePrometheus,
		MockProviderSet,
		ServiceProviderSet,
		DaoProviderSet,
		wire.Struct(new(Injector), "*"),
	)
	return new(Injector), nil
}
