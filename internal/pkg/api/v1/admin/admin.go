package admin

import (
	"data-collection-hub-server/internal/pkg/api/v1/admin/mods"
)

type Admin struct {
	*mods.DataAuditApi
	*mods.StatisticApi
	*mods.UserApi
	*mods.NoticeApi
	*mods.DocumentationApi
	*mods.LogsApi
}

func New(
	dataAuditApi *mods.DataAuditApi, statisticApi *mods.StatisticApi, userApi *mods.UserApi, noticeApi *mods.NoticeApi,
	documentationApi *mods.DocumentationApi, logsApi *mods.LogsApi,
) Admin {
	return Admin{
		DataAuditApi:     dataAuditApi,
		StatisticApi:     statisticApi,
		UserApi:          userApi,
		NoticeApi:        noticeApi,
		DocumentationApi: documentationApi,
		LogsApi:          logsApi,
	}
}
