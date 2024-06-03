package common

import (
	"data-collection-hub-server/internal/pkg/api/v1/common/mods"
)

type Common struct {
	AuthApi          *mods.AuthApi
	ProfileApi       *mods.ProfileApi
	DocumentationApi *mods.DocumentationApi
	NoticeApi        *mods.NoticeApi
	IdempotencyApi   *mods.IdempotencyApi
}
