package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	app "github.com/becosuke/tasks-api/application/grpc/server"
	"github.com/becosuke/tasks-api/config"
)

func main() {
	conf := config.GetConfig()

	listen, err := net.Listen("tcp", conf.GrpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	app.Register(server)

	go func() {
		server.Serve(listen)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	server.GracefulStop()
}
