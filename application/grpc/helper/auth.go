package helper

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"

	service "github.com/becosuke/tasks-api/domain/service/auth"
)

func GetUserId(ctx context.Context) (uint64, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return 0, false
	}

	authorization := md.Get("Authorization")
	if len(authorization) == 0 {
		return 0, false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	userId, ok := service.GetUserId(token)
	if ok == false {
		return 0, false
	}

	return userId, true
}
