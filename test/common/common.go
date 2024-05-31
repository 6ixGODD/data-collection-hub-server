package common

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/utils/crypt"
	"data-collection-hub-server/test/wire"
	"github.com/spf13/viper"
)

func Setup() error {
	cfg := config.New()
	viper.SetConfigFile("../../configs/test/server.test.yml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}
	injector, err := wire.InitializeTestInjector(context.Background(), cfg, 1000)
	if err != nil {
		return err
	}
	wire.SetInjector(injector)
	return nil
}

func Teardown() error {
	injector := wire.GetInjector()

	_, _ = injector.UserDao.DeleteUserList(injector.Ctx, nil, nil, nil, nil, nil, nil, nil, nil)
	_, _ = injector.InstructionDataDao.DeleteInstructionDataList(injector.Ctx, nil, nil, nil, nil, nil, nil, nil)
	_, _ = injector.NoticeDao.DeleteNoticeList(injector.Ctx, nil, nil, nil, nil, nil)
	_, _ = injector.DocumentationDao.DeleteDocumentationList(injector.Ctx, nil, nil, nil, nil)
	_, _ = injector.LoginLogDao.DeleteLoginLogList(injector.Ctx, nil, nil, nil, nil, nil)
	_, _ = injector.OperationLogDao.DeleteOperationLogList(injector.Ctx, nil, nil, nil, nil, nil, nil, nil, nil)
	var (
		username    = "Admin"
		password, _ = crypt.Hash("Admin@123")
		email       = "6goddddddd@gmail.com"
		role        = config.UserRoleAdmin
		org         = "Org"
	)
	_, _ = injector.UserDao.InsertUser(injector.Ctx, username, email, password, role, org)
	err := injector.Cache.Flush(injector.Ctx, nil)
	if err != nil {
		return err
	}
	err = injector.Mongo.Close(injector.Ctx)
	if err != nil {
		return err
	}
	err = injector.Redis.Close()
	if err != nil {
		return err
	}
	return nil
}
