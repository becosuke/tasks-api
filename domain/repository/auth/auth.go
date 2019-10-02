package auth

import (
	entity "github.com/becosuke/tasks-api/domain/entity/auth"
	"github.com/becosuke/tasks-api/infrastructure/cache"
)

func FetchUserIdByToken(token string) (uint64, bool) {
	if len(token) == 0 {
		return 0, false
	}

	cacheKey := entity.GetAccessTokenCacheKey(token)
	res, ok := cache.GetSharedNumber(cacheKey)
	if !ok {
		return 0, false
	}
	return res, true
}
