package dao_test

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	userID            primitive.ObjectID
	username          string
	email             string
	instructionDataID primitive.ObjectID
	noticeID          primitive.ObjectID
	documentID        primitive.ObjectID
	operationLogID    primitive.ObjectID
	loginLogID        primitive.ObjectID
)

//
// func TestMain(m *testing.M) {
// 	config.New().CacheConfig.RedisConfig.Password = "root"
// 	config.New().CasbinConfig.ModelPath = "../../configs/casbin_model.test.conf"
// 	config.New().ZapConfig.Level = "warn"
// 	config.New().ZapConfig.DisableStacktrace = true
// 	config.New().ZapConfig.DisableCaller = true
// 	config.New().MongoConfig.Database = "data-collection-hub-test"
// 	testApp, err = wire.InitializeApp(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	daoCore, err := dao.NewCore(testApp.Ctx, testApp.Mongo, testApp.Zap, testApp.Config)
// 	if err != nil {
// 		panic(err)
// 	}
// 	cache = dao.NewCache(testApp.Redis, testApp.Config)
// 	userDao, err = mods.NewUserDao(testApp.Ctx, daoCore, cache)
// 	if err != nil {
// 		panic(err)
// 	}
// 	instructionDataDao, err = mods.NewInstructionDataDao(testApp.Ctx, daoCore, userDao)
// 	if err != nil {
// 		panic(err)
// 	}
// 	noticeDao, err = mods.NewNoticeDao(testApp.Ctx, daoCore, cache)
// 	if err != nil {
// 		panic(err)
// 	}
// 	documentationDao, err = mods.NewDocumentationDao(testApp.Ctx, daoCore, cache)
// 	if err != nil {
// 		panic(err)
// 	}
// 	operationLogDao, err = mods.NewOperationLogDao(testApp.Ctx, daoCore, cache, userDao)
// 	if err != nil {
// 		panic(err)
// 	}
// 	loginLogDao, err = mods.NewLoginLogDao(testApp.Ctx, daoCore, cache, userDao)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	mockUser = mock.NewUserDaoMockWithRandomData(1000, userDao)
// 	mockInstructionData = mock.NewInstructionDataDaoMockWithRandomData(1000, mockUser, instructionDataDao)
// 	mockNotice = mock.NewNoticeDaoMockWithRandomData(1000, noticeDao)
// 	mockDocumentation = mock.NewDocumentationDaoMockWithRandomData(1000, documentationDao)
// 	mockOperationLog = mock.NewOperationLogDaoMockWithRandomData(
// 		1000, operationLogDao, mockUser, mockInstructionData, mockNotice, mockDocumentation,
// 	)
// 	mockLoginLog = mock.NewLoginLogDaoMockWithRandomData(1000, loginLogDao, mockUser)
//
// 	code := m.Run()
//
// 	mockUser.Delete()
// 	mockInstructionData.Delete()
// 	mockNotice.Delete()
// 	mockDocumentation.Delete()
// 	mockOperationLog.Delete()
// 	mockLoginLog.Delete()
//
// 	_, _ = userDao.DeleteUserList(testApp.Ctx, nil, nil, nil, nil, nil, nil, nil, nil)
// 	_, _ = instructionDataDao.DeleteInstructionDataList(testApp.Ctx, nil, nil, nil, nil, nil, nil, nil)
// 	_, _ = noticeDao.DeleteNoticeList(testApp.Ctx, nil, nil, nil, nil, nil)
// 	_, _ = documentationDao.DeleteDocumentationList(testApp.Ctx, nil, nil, nil, nil)
// 	_, _ = operationLogDao.DeleteOperationLogList(testApp.Ctx, nil, nil, nil, nil, nil, nil, nil, nil)
// 	_, _ = loginLogDao.DeleteLoginLogList(testApp.Ctx, nil, nil, nil, nil, nil)
// 	err = testApp.Mongo.Close(testApp.Ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = testApp.Redis.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	os.Exit(code)
// }
