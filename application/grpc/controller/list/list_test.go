package list

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"
	"os"
	"testing"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/becosuke/tasks-api/config"
	entity "github.com/becosuke/tasks-api/domain/entity/list"
	service "github.com/becosuke/tasks-api/domain/service/list"
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
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			grpcvalidator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpcvalidator.UnaryServerInterceptor(),
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

func create(title string) (*entity.Document, error) {
	if !config.IsLocal() {
		return nil, errors.New("skip test")
	}

	document, err := service.Create(title)
	if err != nil {
		return &entity.Document{}, err
	}

	bs, _ := json.Marshal(document.Message())
	log.Print(string(bs))
	return document, nil
}

func TestCreate(t *testing.T) {
	if !config.IsLocal() {
		t.Log("skip test")
		return
	}

	ctx := context.Background()
	req := &pbmessage.CreateRequest{Title: "created"}
	res, err := client.Create(ctx, req)
	if err != nil {
		t.Error(err)
	}

	bs, _ := json.Marshal(res)
	log.Print(string(bs))
}

func TestUpdate(t *testing.T) {
	if !config.IsLocal() {
		t.Log("skip test")
		return
	}

	created, err := create("update")
	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()
	req := &pbmessage.UpdateRequest{Id: created.Id, Title: "updated"}
	res, err := client.Update(ctx, req)
	if err != nil {
		t.Error(err)
	}

	bs, _ := json.Marshal(res)
	log.Print(string(bs))
}

func TestDelete(t *testing.T) {
	if !config.IsLocal() {
		t.Log("skip test")
		return
	}

	created, err := create("delete")
	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()
	req := &pbmessage.DeleteRequest{Id: created.Id}
	res, err := client.Delete(ctx, req)
	if err != nil {
		t.Error(err)
	}

	bs, _ := json.Marshal(res)
	log.Print(string(bs))
}

func TestGetDocument(t *testing.T) {
	ctx := context.Background()
	req := &pbmessage.GetDocumentRequest{Id: 1}
	res, err := client.GetDocument(ctx, req)
	if err != nil {
		t.Error(err)
	}

	bs, _ := json.Marshal(res)
	log.Print(string(bs))
}

func TestGetDocuments(t *testing.T) {
	ctx := context.Background()
	req := &pbmessage.GetDocumentsRequest{Ids: []uint64{1, 2, 3}}
	res, err := client.GetDocuments(ctx, req)
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
