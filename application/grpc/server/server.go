package server

import (
	"google.golang.org/grpc"

	"github.com/becosuke/tasks-api/application/grpc/server/context"
	"github.com/becosuke/tasks-api/application/grpc/server/list"
	"github.com/becosuke/tasks-api/application/grpc/server/task"
)

func Register(server *grpc.Server) {
	list.Register(server)
	context.Register(server)
	task.Register(server)
}
