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
