package mods

import (
	"data-collection-hub-server/internal/pkg/dao"
)

type IdempotencyMiddleware struct {
	Cache *dao.Cache
}
