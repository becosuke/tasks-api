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
	var err error
	conf := config.GetConfig()

	ctx := context.Background()
	r.ctx, r.Cancel = context.WithCancel(ctx)

	r.mux = runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(headerMatcher))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	r.conn, err = grpc.Dial(conf.GrpcAddr, opts...)
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
	var err error
	conf := config.GetConfig()

	if err != nil {
		if cerr := r.conn.Close(); cerr != nil {
			grpclog.Infof("Failed to close conn to %s: %v", conf.GrpcAddr, cerr)
		}
		return
	}
	go func() {
		<-r.ctx.Done()
		if cerr := r.conn.Close(); cerr != nil {
			grpclog.Infof("Failed to close conn to %s: %v", conf.GrpcAddr, cerr)
		}
	}()
}

func (r *Router) Run() error {
	conf := config.GetConfig()

	return http.ListenAndServe(conf.RestAddr, r.mux)
}

func headerMatcher(headerName string) (string, bool) {
	mdName := "x-device-type"
	if headerName == "X-Device-Type" {
		return mdName, true
	} else {
		return mdName, false
	}
}
