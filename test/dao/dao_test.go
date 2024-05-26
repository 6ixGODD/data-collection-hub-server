package dao_test

import (
	"context"
	"os"
	"testing"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	"data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/wire"
	"data-collection-hub-server/test/dao/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	cache *dao.Cache
)

// User Dao Test Variables
var (
	userDaoCtx context.Context
	userDao    mods.UserDao
	userID     primitive.ObjectID
	username   string
	email      string
	mockUser   *mock.UserDaoMock
	err        error
)

// Instruction Data Dao Test Variables
var (
	instructionDataDaoCtx context.Context
	instructionDataDao    mods.InstructionDataDao
	instructionDataID     primitive.ObjectID
	mockInstructionData   *mock.InstructionDataDaoMock
)

// Notice Dao Test Variables
var (
	noticeDaoCtx context.Context
	noticeDao    mods.NoticeDao
	noticeID     primitive.ObjectID
	mockNotice   *mock.NoticeDaoMock
)

// Documentation Dao Test Variables
var (
	documentationDaoCtx context.Context
	documentationDao    mods.DocumentationDao
	documentID          primitive.ObjectID
	mockDocumentation   *mock.DocumentationDaoMock
)

// Operation Log Dao Test Variables
var (
	operationLogDaoCtx context.Context
	operationLogDao    mods.OperationLogDao
	operationLogID     primitive.ObjectID
	mockOperationLog   *mock.OperationLogDaoMock
)

// Login Log Dao Test Variables
var (
	loginLogDaoCtx context.Context
	loginLogDao    mods.LoginLogDao
	loginLogID     primitive.ObjectID
	mockLoginLog   *mock.LoginLogDaoMock
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	cfg := config.New()
	cfg.MongoConfig.Database = "data-collection-hub"
	cfg.CacheConfig.RedisConfig.Password = "root"
	cfg.ZapConfig.Level = "warn"
	cfg.ZapConfig.DisableStacktrace = true
	cfg.ZapConfig.DisableCaller = true
	mongo, err := wire.InitializeMongo(ctx, cfg)
	if err != nil {
		panic(err)
	}
	redis, err := wire.InitializeRedis(ctx, cfg)
	if err != nil {
		panic(err)
	}
	zap, err := wire.InitializeZap(cfg)
	if err != nil {
		panic(err)
	}
	core, err := dao.NewCore(ctx, mongo, zap, cfg)
	if err != nil {
		panic(err)
	}
	cache = dao.NewCache(redis, cfg)

	// User Dao Test Setup
	userDaoCtx = ctx
	userDao, err = mods.NewUserDao(ctx, core, cache)
	if err != nil {
		panic(err)
	}
	mockUser = mock.NewUserDaoMockWithRandomData(1000, userDao)

	// Instruction Data Dao Test Setup
	instructionDataDaoCtx = ctx
	instructionDataDao, err = mods.NewInstructionDataDao(ctx, core, userDao)
	if err != nil {
		panic(err)
	}
	mockInstructionData = mock.NewInstructionDataDaoMockWithRandomData(1000, mockUser, instructionDataDao)

	// Notice Dao Test Setup
	noticeDaoCtx = ctx
	noticeDao, err = mods.NewNoticeDao(ctx, core, cache)
	if err != nil {
		panic(err)
	}
	mockNotice = mock.NewNoticeDaoMockWithRandomData(1000, noticeDao)

	// Documentation Dao Test Setup
	documentationDaoCtx = ctx
	documentationDao, err = mods.NewDocumentationDao(ctx, core, cache)
	if err != nil {
		panic(err)
	}
	mockDocumentation = mock.NewDocumentationDaoMockWithRandomData(1000, documentationDao)

	// Operation Log Dao Test Setup
	operationLogDaoCtx = ctx
	operationLogDao, err = mods.NewOperationLogDao(ctx, core, cache, userDao)
	if err != nil {
		panic(err)
	}
	mockOperationLog = mock.NewOperationLogDaoMockWithRandomData(
		1000, operationLogDao, *mockUser, *mockInstructionData, *mockNotice, *mockDocumentation,
	)

	// Login Log Dao Test Setup
	loginLogDaoCtx = ctx
	loginLogDao, err = mods.NewLoginLogDao(ctx, core, cache, userDao)
	if err != nil {
		panic(err)
	}
	mockLoginLog = mock.NewLoginLogDaoMockWithRandomData(1000, loginLogDao, *mockUser)

	code := m.Run()

	mockUser.Delete()
	mockInstructionData.Delete()
	mockNotice.Delete()
	mockDocumentation.Delete()
	mockOperationLog.Delete()
	mockLoginLog.Delete()

	_, _ = userDao.DeleteUserList(ctx, nil, nil, nil, nil, nil, nil, nil, nil)
	_, _ = instructionDataDao.DeleteInstructionDataList(ctx, nil, nil, nil, nil, nil, nil, nil)
	_, _ = noticeDao.DeleteNoticeList(ctx, nil, nil, nil, nil, nil)
	_, _ = documentationDao.DeleteDocumentationList(ctx, nil, nil, nil, nil)
	_, _ = operationLogDao.DeleteOperationLogList(ctx, nil, nil, nil, nil, nil, nil, nil, nil)
	_, _ = loginLogDao.DeleteLoginLogList(ctx, nil, nil, nil, nil, nil)
	err = mongo.Close(ctx)
	if err != nil {
		panic(err)
	}
	err = redis.Close()
	if err != nil {
		panic(err)
	}
	os.Exit(code)
}
