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

func (s *Server) GetDocument(ctx context.Context, request *common.Id) (*list.Document, error) {
	if val, err := service.GetDocument(request.Id); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return val.Message(), nil
	}
}

func (s *Server) GetDocuments(request *common.Ids, stream stub.List_GetDocumentsServer) error {
	if vals, err := service.GetDocuments(request.Ids); err != nil {
		log.Printf("%+v", err)
		return err
	} else {
		for _, val := range vals {
			if err := stream.Send(val.Message()); err != nil {
				log.Printf("%+v", err)
				return err
			}
		}
	}
	return nil
}

func (s *Server) GetDocumentsAll(request *common.Pagination, stream stub.List_GetDocumentsAllServer) error {
	if vals, err := service.GetDocumentsAll(request.Limit, request.Offset); err != nil {
		log.Printf("%+v", err)
		return err
	} else {
		for _, val := range vals {
			if err := stream.Send(val.Message()); err != nil {
				log.Printf("%+v", err)
				return err
			}
		}
	}
	return nil
}

func (s *Server) GetCountAll(ctx context.Context, request *common.Empty) (*common.Count, error) {
	if res, err := service.GetCountAll(); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return &common.Count{Count: res}, nil
	}
}

func (s *Server) Create(ctx context.Context, request *list.CreateRequest) (*list.Document, error) {
	if val, err := service.Create(request.Title); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return val.Message(), nil
	}
}

func (s *Server) Update(ctx context.Context, request *list.UpdateRequest) (*list.Document, error) {
	if val, err := service.Update(request.Id, request.Title); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return val.Message(), nil
	}
}

func (s *Server) Delete(ctx context.Context, request *list.DeleteRequest) (*list.Document, error) {
	if val, err := service.Delete(request.Id); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return val.Message(), nil
	}
}
