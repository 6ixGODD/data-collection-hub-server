package service_test

import (
	"context"
	"os"
	"testing"

	"data-collection-hub-server/internal/app"
	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	"data-collection-hub-server/internal/pkg/dao/mods"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	commonservice "data-collection-hub-server/internal/pkg/service/common/mods"
	sysservice "data-collection-hub-server/internal/pkg/service/sys/mods"
	userservice "data-collection-hub-server/internal/pkg/service/user/mods"
	"data-collection-hub-server/internal/pkg/wire"
	"data-collection-hub-server/test/mock"
)

var (
	testApp *app.App
	err     error
)

var (
	instructionDataDao mods.InstructionDataDao
	noticeDao          mods.NoticeDao
	documentationDao   mods.DocumentationDao
	userDao            mods.UserDao
	loginLogDao        mods.LoginLogDao
	operationLogDao    mods.OperationLogDao
)

var (
	mockInstructionData *mock.InstructionDataDaoMock
	mockNotice          *mock.NoticeDaoMock
	mockDocumentation   *mock.DocumentationDaoMock
	mockUser            *mock.UserDaoMock
	mockLoginLog        *mock.LoginLogDaoMock
	mockOperationLog    *mock.OperationLogDaoMock
)

var (
	adminDataAuditService     adminservice.DataAuditService
	adminDocumentationService adminservice.DocumentationService
	adminNoticeService        adminservice.NoticeService
	adminLogsService          adminservice.LogsService
	adminStatisticService     adminservice.StatisticService
	adminUserService          adminservice.UserService

	commonAuthService          commonservice.AuthService
	commonIdempotencyService   commonservice.IdempotencyService
	commonDocumentationService commonservice.DocumentationService
	commonNoticeService        commonservice.NoticeService
	commonProfileService       commonservice.ProfileService

	sysLogsService sysservice.LogsService

	userDatasetService   userservice.DatasetService
	userStatisticService userservice.StatisticService
)

func TestMain(t *testing.M) {
	config.New().CacheConfig.RedisConfig.Password = "root"
	config.New().CasbinConfig.ModelPath = "../../configs/casbin_model.test.conf"
	config.New().ZapConfig.Level = "warn"
	config.New().ZapConfig.DisableStacktrace = true
	config.New().ZapConfig.DisableCaller = true
	config.New().MongoConfig.Database = "data-collection-hub-test"
	testApp, err = wire.InitializeApp(context.Background())
	if err != nil {
		panic(err)
	}

	adminDataAuditService = testApp.Router.RouterV1.ApiV1.AdminApi.DataAuditApi.DataAuditService
	adminDocumentationService = testApp.Router.RouterV1.ApiV1.AdminApi.DocumentationApi.DocumentationService
	adminNoticeService = testApp.Router.RouterV1.ApiV1.AdminApi.NoticeApi.NoticeService
	adminLogsService = testApp.Router.RouterV1.ApiV1.AdminApi.LogsApi.LogsService
	adminStatisticService = testApp.Router.RouterV1.ApiV1.AdminApi.StatisticApi.StatisticService
	adminUserService = testApp.Router.RouterV1.ApiV1.AdminApi.UserApi.UserService
	commonAuthService = testApp.Router.RouterV1.ApiV1.CommonApi.AuthApi.AuthService
	commonIdempotencyService = testApp.Middleware.IdempotencyMiddleware.IdempotencyService
	commonDocumentationService = testApp.Router.RouterV1.ApiV1.CommonApi.DocumentationApi.DocumentationService
	commonNoticeService = testApp.Router.RouterV1.ApiV1.CommonApi.NoticeApi.NoticeService
	commonProfileService = testApp.Router.RouterV1.ApiV1.CommonApi.ProfileApi.ProfileService
	sysLogsService = testApp.Router.RouterV1.ApiV1.CommonApi.AuthApi.LogsService
	userDatasetService = testApp.Router.RouterV1.ApiV1.UserApi.DatasetApi.DatasetService
	userStatisticService = testApp.Router.RouterV1.ApiV1.UserApi.StatisticApi.StatisticService

	daoCore, err := dao.NewCore(testApp.Ctx, testApp.Mongo, testApp.Zap, testApp.Config)
	if err != nil {
		panic(err)
	}
	cache := dao.NewCache(testApp.Redis, testApp.Config)
	userDao, err = mods.NewUserDao(testApp.Ctx, daoCore, cache)
	if err != nil {
		panic(err)
	}
	instructionDataDao, err = mods.NewInstructionDataDao(testApp.Ctx, daoCore, userDao)
	if err != nil {
		panic(err)
	}
	noticeDao, err = mods.NewNoticeDao(testApp.Ctx, daoCore, cache)
	if err != nil {
		panic(err)
	}
	documentationDao, err = mods.NewDocumentationDao(testApp.Ctx, daoCore, cache)
	if err != nil {
		panic(err)
	}
	loginLogDao, err = mods.NewLoginLogDao(testApp.Ctx, daoCore, cache, userDao)
	if err != nil {
		panic(err)
	}
	operationLogDao, err = mods.NewOperationLogDao(testApp.Ctx, daoCore, cache, userDao)
	if err != nil {
		panic(err)
	}

	mockUser = mock.NewUserDaoMockWithRandomData(1000, userDao)
	mockInstructionData = mock.NewInstructionDataDaoMockWithRandomData(1000, mockUser, instructionDataDao)
	mockNotice = mock.NewNoticeDaoMockWithRandomData(1000, noticeDao)
	mockDocumentation = mock.NewDocumentationDaoMockWithRandomData(1000, documentationDao)
	mockLoginLog = mock.NewLoginLogDaoMockWithRandomData(1000, loginLogDao, *mockUser)
	mockOperationLog = mock.NewOperationLogDaoMockWithRandomData(
		1000, operationLogDao, *mockUser, *mockInstructionData, *mockNotice, *mockDocumentation,
	)

	code := t.Run()

	mockUser.Delete()
	mockInstructionData.Delete()
	mockNotice.Delete()
	mockDocumentation.Delete()
	mockOperationLog.Delete()
	mockLoginLog.Delete()

	_, _ = userDao.DeleteUserList(testApp.Ctx, nil, nil, nil, nil, nil, nil, nil, nil)
	_, _ = instructionDataDao.DeleteInstructionDataList(testApp.Ctx, nil, nil, nil, nil, nil, nil, nil)
	_, _ = noticeDao.DeleteNoticeList(testApp.Ctx, nil, nil, nil, nil, nil)
	_, _ = documentationDao.DeleteDocumentationList(testApp.Ctx, nil, nil, nil, nil)
	_, _ = operationLogDao.DeleteOperationLogList(testApp.Ctx, nil, nil, nil, nil, nil, nil, nil, nil)
	_, _ = loginLogDao.DeleteLoginLogList(testApp.Ctx, nil, nil, nil, nil, nil)
	err = testApp.Mongo.Close(testApp.Ctx)
	if err != nil {
		panic(err)
	}
	err = testApp.Redis.Close()
	if err != nil {
		panic(err)
	}
	os.Exit(code)
}
