package list

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/becosuke/tasks-api/application/rest/helper/router"
	pbservice "github.com/becosuke/tasks-api/protogen/service/list"
)

func registerer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	var err error

	if err = pbservice.RegisterListHandler(ctx, mux, conn); err != nil {
		return err
	}

	return nil
}

func NewRouter() *router.Router {
	return router.NewRouter(registerer)
}
