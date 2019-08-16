package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"

	"github.com/becosuke/tasks-api/application/grpc/controller"
	"github.com/becosuke/tasks-api/config"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {
	listen, err := net.Listen("tcp", config.GetConfig().GrpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)

	controller.Register(server)

	go func() {
		server.Serve(listen)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	server.GracefulStop()
}
