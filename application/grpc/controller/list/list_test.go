package list

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"testing"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pbmessage "github.com/becosuke/tasks-api/protogen/message/list"
	pbservice "github.com/becosuke/tasks-api/protogen/service/list"
)

var server *grpc.Server
var client pbservice.ListClient

func TestMain(m *testing.M) {
	log.SetFlags(log.Ldate | log.Ltime)
	if err := setup(); err != nil {
		log.Fatal(err)
	}
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() error {
	listen := bufconn.Listen(1024 * 1024)
	server = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)
	Register(server)

	go func() {
		if err := server.Serve(listen); err != nil {
			log.Fatal(err)
			return
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufcon", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return listen.Dial() }), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client = pbservice.NewListClient(conn)

	return nil
}

func teardown() {
	server.GracefulStop()
}

func TestGetCountAll(t *testing.T) {
	ctx := context.Background()
	req := &pbmessage.GetCountAllRequest{}
	res, err := client.GetCountAll(ctx, req)
	if err != nil {
		t.Error(err)
	}

	bs, _ := json.Marshal(res)
	log.Print(string(bs))
}

func TestGetDocumentsAll(t *testing.T) {
	ctx := context.Background()
	req := &pbmessage.GetDocumentsAllRequest{Limit: 3, Offset: 0}
	res, err := client.GetDocumentsAll(ctx, req)
	if err != nil {
		t.Error(err)
	}

	bs, _ := json.Marshal(res)
	log.Print(string(bs))
}
