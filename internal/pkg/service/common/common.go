package common

import (
	"data-collection-hub-server/internal/pkg/service/common/mods"
)

type Common struct {
	AuthService          mods.AuthService
	DocumentationService mods.DocumentationService
	NoticeService        mods.NoticeService
	ProfileService       mods.ProfileService
}

func New(
	authService mods.AuthService, documentationService mods.DocumentationService,
	noticeService mods.NoticeService, profileService mods.ProfileService,
) *Common {
	return &Common{
		AuthService:          authService,
		DocumentationService: documentationService,
		NoticeService:        noticeService,
		ProfileService:       profileService,
	}
}
