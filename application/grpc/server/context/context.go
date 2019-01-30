package context

import (
	"context"
	"log"

	"google.golang.org/grpc"

	service "github.com/becosuke/tasks-api/domain/service/context"
	"github.com/becosuke/tasks-api/protogen/message/common"
	message "github.com/becosuke/tasks-api/protogen/message/context"
	stub "github.com/becosuke/tasks-api/protogen/service/context"
)

type Server struct{}

func Register(grpcServer *grpc.Server) {
	stub.RegisterContextServer(grpcServer, &Server{})
}

func (s *Server) GetDocument(ctx context.Context, request *common.Id) (*message.Document, error) {
	val, err := service.GetDocument(request.Id)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	return val.Message(), nil
}

func (s *Server) GetDocuments(request *common.Ids, stream stub.Context_GetDocumentsServer) error {
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

func (s *Server) GetDocumentsAll(request *common.Pagination, stream stub.Context_GetDocumentsAllServer) error {
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
