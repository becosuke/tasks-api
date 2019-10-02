package context

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/becosuke/tasks-api/application/rest/helper/base"
	pbservice "github.com/becosuke/tasks-api/protogen/service/context"
)

func registerer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	var err error

	if err = pbservice.RegisterContextHandler(ctx, mux, conn); err != nil {
		return err
	}

	return nil
}

func NewRouter() *base.Router {
	return base.NewRouter(registerer)
}
