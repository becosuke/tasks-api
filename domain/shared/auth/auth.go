package auth

import (
	"github.com/becosuke/tasks-api/domain/repository/auth"
)

func GetUserId(token string) (uint64, bool) {
	return auth.FetchUserIdByToken(token)
}
