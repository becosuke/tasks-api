package list

import (
	"context"
	"log"

	"google.golang.org/grpc"

	service "github.com/becosuke/tasks-api/domain/service/list"
	"github.com/becosuke/tasks-api/protogen/message/common"
	"github.com/becosuke/tasks-api/protogen/message/list"
	stub "github.com/becosuke/tasks-api/protogen/service/list"
)

type Server struct{}

func Register(grpcServer *grpc.Server) {
	stub.RegisterListServer(grpcServer, &Server{})
}

func (s *Server) Create(ctx context.Context, request *list.CreateRequest) (*list.Document, error) {
	val, err := service.Create(request.Title)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	return val.Message(), nil
}

func (s *Server) Update(ctx context.Context, request *list.UpdateRequest) (*list.Document, error) {
	val, err := service.Update(request.Id, request.Title)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	return val.Message(), nil
}

func (s *Server) Delete(ctx context.Context, request *list.DeleteRequest) (*list.Document, error) {
	val, err := service.Delete(request.Id)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	return val.Message(), nil
}

func (s *Server) GetDocument(ctx context.Context, request *common.Id) (*list.Document, error) {
	val, err := service.GetDocument(request.Id)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	return val.Message(), nil
}

func (s *Server) GetDocuments(request *common.Ids, stream stub.List_GetDocumentsServer) error {
	vals, err := service.GetDocuments(request.Ids)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	for _, val := range vals {
		if err := stream.Send(val.Message()); err != nil {
			log.Printf("%+v", err)
			return err
		}
	}

	return nil
}

func (s *Server) GetDocumentsAll(request *common.Pagination, stream stub.List_GetDocumentsAllServer) error {
	vals, err := service.GetDocumentsAll(request.Limit, request.Offset)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	for _, val := range vals {
		if err := stream.Send(val.Message()); err != nil {
			log.Printf("%+v", err)
			return err
		}
	}

	return nil
}

func (s *Server) GetCountAll(ctx context.Context, request *common.Empty) (*common.Count, error) {
	val, err := service.GetCountAll()
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	return val.Message(), nil
}
