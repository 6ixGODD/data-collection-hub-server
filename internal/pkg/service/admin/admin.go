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

// New creates a new Admin service.
func New(
	dataAuditService mods.DataAuditService, documentationService mods.DocumentationService,
	logsService mods.LogsService, noticeService mods.NoticeService, statisticService mods.StatisticService,
	userService mods.UserService,
) *Admin {
	return &Admin{
		DataAuditService:     dataAuditService,
		DocumentationService: documentationService,
		LogsService:          logsService,
		NoticeService:        noticeService,
		StatisticService:     statisticService,
		UserService:          userService,
	}
}
