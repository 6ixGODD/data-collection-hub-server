package mods

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/utils/common"
)

type IdempotencyService interface {
	GenerateIdempotencyToken(ctx context.Context) (string, error)
	CheckIdempotencyToken(ctx context.Context, token string) (bool, error)
}

type idempotencyServiceImpl struct {
	core  *service.Core
	cache *dao.Cache
}

func NewIdempotencyService(core *service.Core, cache *dao.Cache) IdempotencyService {
	return &idempotencyServiceImpl{
		core:  core,
		cache: cache,
	}
}

func (s *idempotencyServiceImpl) GenerateIdempotencyToken(ctx context.Context) (string, error) {
	token, err := common.GenerateUUID()
	if err != nil {
		return "", err
	}
	if err = s.cache.Set(ctx, token, config.CacheTrue, &s.core.Config.IdempotencyConfig.TTL); err != nil {
		return "", err
	}
	return token, nil
}

func (s *idempotencyServiceImpl) CheckIdempotencyToken(ctx context.Context, token string) (bool, error) {
	result, err := s.cache.Get(ctx, token)
	if err != nil {
		return false, err
	}
	if err = s.cache.Delete(ctx, token); err != nil {
		return false, err
	}
	return result != nil, nil
}
