package common

import (
	"data-collection-hub-server/internal/pkg/api/v1/common/mods"
)

type Common struct {
	*mods.AuthApi
	*mods.ProfileApi
	*mods.DocumentationApi
	*mods.NoticeApi
}

func New(
	authApi *mods.AuthApi, profileApi *mods.ProfileApi, documentationApi *mods.DocumentationApi,
	noticeApi *mods.NoticeApi,
) *Common {
	return &Common{
		AuthApi:          authApi,
		ProfileApi:       profileApi,
		DocumentationApi: documentationApi,
		NoticeApi:        noticeApi,
	}
}
