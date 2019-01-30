package helper

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"

	service "github.com/becosuke/tasks-api/domain/service/auth"
)

func GetUserID(ctx context.Context) (uint64, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return 0, false
	}

	authorization := md.Get("Authorization")
	if len(authorization) == 0 {
		return 0, false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	userID, ok := service.GetUserID(token)
	if ok == false {
		return 0, false
	}

	return userID, true
}
