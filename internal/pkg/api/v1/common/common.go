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
