package admin

import (
	"data-collection-hub-server/internal/pkg/api/v1/admin/mods"
)

type Admin struct {
	DataAuditApi     *mods.DataAuditApi
	StatisticApi     *mods.StatisticApi
	UserApi          *mods.UserApi
	NoticeApi        *mods.NoticeApi
	DocumentationApi *mods.DocumentationApi
	LogsApi          *mods.LogsApi
}
