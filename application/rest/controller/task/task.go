package task

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/becosuke/tasks-api/application/rest/helper/router"
	stub "github.com/becosuke/tasks-api/protogen/service/task"
)

func registerer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	var err error

	if err = stub.RegisterTaskHandler(ctx, mux, conn); err != nil {
		return err
	}

	return nil
}

func NewRouter() *router.Router {
	return router.NewRouter(registerer)
}
