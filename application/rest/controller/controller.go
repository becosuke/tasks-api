package controller

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/becosuke/tasks-api/application/rest/helper/base"
	pbcont "github.com/becosuke/tasks-api/protogen/service/context"
	pblist "github.com/becosuke/tasks-api/protogen/service/list"
	pbtask "github.com/becosuke/tasks-api/protogen/service/task"
)

func registerer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	var err error

	if err = pblist.RegisterListHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err = pbcont.RegisterContextHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err = pbtask.RegisterTaskHandler(ctx, mux, conn); err != nil {
		return err
	}

	return nil
}

func NewRouter() *base.Router {
	return base.NewRouter(registerer)
}
