package admin

import (
	"data-collection-hub-server/internal/pkg/service/admin/mods"
)

type Admin struct {
	DataAuditService     mods.DataAuditService
	DocumentationService mods.DocumentationService
	LogsService          mods.LogsService
	NoticeService        mods.NoticeService
	StatisticService     mods.StatisticService
	UserService          mods.UserService
}
