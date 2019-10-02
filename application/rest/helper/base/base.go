package base

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/becosuke/tasks-api/config"
)

type Router struct {
	ctx      context.Context
	mux      *runtime.ServeMux
	conn     *grpc.ClientConn
	Register func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
	Cancel   func()
}

func NewRouter(f func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error) *Router {
	return &Router{Register: f}
}

func (r *Router) GetConn() *grpc.ClientConn {
	return r.conn
}

func (r *Router) GetMux() *runtime.ServeMux {
	return r.mux
}

func (r *Router) Setup() error {
	ctx := context.Background()
	r.ctx, r.Cancel = context.WithCancel(ctx)

	r.mux = runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(func (headerName string) (string, bool) {
		mdName := "x-device-type"
		if headerName == "X-Device-Type" {
			return mdName, true
		}

		return mdName, false
	}))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	var err error
	r.conn, err = grpc.Dial(config.GetConfig().GrpcAddr, opts...)
	if err != nil {
		return err
	}

	err = r.Register(r.ctx, r.mux, r.conn)
	if err != nil {
		return err
	}

	return nil
}

func (r *Router) Close() {
	go func() {
		<-r.ctx.Done()
		if err := r.conn.Close(); err != nil {
			grpclog.Infof("Failed to close conn to %s: %v", config.GetConfig().GrpcAddr, err)
		}
	}()
}

func (r *Router) Run() error {
	conf := config.GetConfig()

	return http.ListenAndServe(conf.RestAddr, r.mux)
}
