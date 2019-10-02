package auth

import (
	"fmt"
)

const AccessTokenCacheKey = "native_access_token_%s"

func GetAccessTokenCacheKey(token string) string {
	return fmt.Sprintf(AccessTokenCacheKey, token)
}
