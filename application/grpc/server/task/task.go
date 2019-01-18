package task

import (
	"context"
	"log"

	"google.golang.org/grpc"

	service "github.com/becosuke/tasks-api/domain/service/task"
	"github.com/becosuke/tasks-api/protogen/message/common"
	"github.com/becosuke/tasks-api/protogen/message/task"
	stub "github.com/becosuke/tasks-api/protogen/service/task"
)

type Server struct{}

func Register(grpcServer *grpc.Server) {
	stub.RegisterTaskServer(grpcServer, &Server{})
}

func (s *Server) Create(ctx context.Context, request *task.CreateRequest) (*task.Document, error) {
	if val, err := service.Create(request.ListId, request.Title); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return val.Message(), nil
	}
}

func (s *Server) Update(ctx context.Context, request *task.UpdateRequest) (*task.Document, error) {
	if val, err := service.Update(request.Id, request.ListId, request.Title); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return val.Message(), nil
	}
}

func (s *Server) Delete(ctx context.Context, request *task.DeleteRequest) (*task.Document, error) {
	if val, err := service.Delete(request.Id); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return val.Message(), nil
	}
}

func (s *Server) GetDocument(ctx context.Context, request *common.Id) (*task.Document, error) {
	if val, err := service.GetDocument(request.Id); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return val.Message(), nil
	}
}

func (s *Server) GetDocuments(request *common.Ids, stream stub.Task_GetDocumentsServer) error {
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

func (s *Server) GetDocumentsAll(request *common.Pagination, stream stub.Task_GetDocumentsAllServer) error {
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
	if val, err := service.GetCountAll(); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return val.Message(), nil
	}
}

func (s *Server) GetDocumentsByList(request *common.IdPagination, stream stub.Task_GetDocumentsByListServer) error {
	if vals, err := service.GetDocumentsByList(request.Id, request.Limit, request.Offset); err != nil {
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

func (s *Server) GetCountByList(ctx context.Context, request *common.Id) (*common.Count, error) {
	if val, err := service.GetCountByList(request.Id); err != nil {
		log.Printf("%+v", err)
		return nil, err
	} else {
		return val.Message(), nil
	}
}
