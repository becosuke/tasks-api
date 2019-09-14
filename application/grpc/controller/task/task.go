package task

import (
	"context"

	"google.golang.org/grpc"

	service "github.com/becosuke/tasks-api/domain/service/task"
	"github.com/becosuke/tasks-api/logger"
	pbmessage "github.com/becosuke/tasks-api/protogen/message/task"
	pbservice "github.com/becosuke/tasks-api/protogen/service/task"
)

type Server struct{}

func Register(grpcServer *grpc.Server) {
	pbservice.RegisterTaskServer(grpcServer, &Server{})
}

func (s *Server) Create(ctx context.Context, request *pbmessage.CreateRequest) (*pbmessage.CreateResponse, error) {
	document, err := service.Create(request.ListId, request.Title)
	if err != nil {
		logger.Error(err)
		return &pbmessage.CreateResponse{}, err
	}

	return &pbmessage.CreateResponse{Result: true, Document: document.Message()}, nil
}

func (s *Server) Update(ctx context.Context, request *pbmessage.UpdateRequest) (*pbmessage.UpdateResponse, error) {
	document, err := service.Update(request.Id, request.Title)
	if err != nil {
		logger.Error(err)
		return &pbmessage.UpdateResponse{}, err
	}

	return &pbmessage.UpdateResponse{Result: true, Document: document.Message()}, nil
}

func (s *Server) Delete(ctx context.Context, request *pbmessage.DeleteRequest) (*pbmessage.DeleteResponse, error) {
	document, err := service.Delete(request.Id)
	if err != nil {
		logger.Error(err)
		return &pbmessage.DeleteResponse{}, err
	}

	return &pbmessage.DeleteResponse{Result: true, Document: document.Message()}, nil
}

func (s *Server) GetDocument(ctx context.Context, request *pbmessage.GetDocumentRequest) (*pbmessage.GetDocumentResponse, error) {
	document, err := service.GetDocument(request.Id)
	if err != nil {
		logger.Error(err)
		return &pbmessage.GetDocumentResponse{}, err
	}

	return &pbmessage.GetDocumentResponse{Document: document.Message()}, nil
}

func (s *Server) GetDocuments(ctx context.Context, request *pbmessage.GetDocumentsRequest) (*pbmessage.GetDocumentsResponse, error) {
	documents, err := service.GetDocuments(request.Ids)
	if err != nil {
		logger.Error(err)
		return &pbmessage.GetDocumentsResponse{}, err
	}

	return &pbmessage.GetDocumentsResponse{Documents: documents.Message()}, nil
}

func (s *Server) GetDocumentsAll(ctx context.Context, request *pbmessage.GetDocumentsAllRequest) (*pbmessage.GetDocumentsAllResponse, error) {
	documents, err := service.GetDocumentsAll(request.Limit, request.Offset)
	if err != nil {
		logger.Error(err)
		return &pbmessage.GetDocumentsAllResponse{}, err
	}

	return &pbmessage.GetDocumentsAllResponse{Documents: documents.Message()}, nil
}

func (s *Server) GetCountAll(ctx context.Context, request *pbmessage.GetCountAllRequest) (*pbmessage.GetCountAllResponse, error) {
	count, err := service.GetCountAll()
	if err != nil {
		logger.Error(err)
		return &pbmessage.GetCountAllResponse{}, err
	}

	return &pbmessage.GetCountAllResponse{Count: count.Message()}, nil
}

func (s *Server) GetDocumentsByList(ctx context.Context, request *pbmessage.GetDocumentsByListRequest) (*pbmessage.GetDocumentsByListResponse, error) {
	documents, err := service.GetDocumentsByList(request.ListId, request.Limit, request.Offset)
	if err != nil {
		logger.Error(err)
		return &pbmessage.GetDocumentsByListResponse{}, err
	}

	return &pbmessage.GetDocumentsByListResponse{Documents: documents.Message()}, nil
}

func (s *Server) GetCountByList(ctx context.Context, request *pbmessage.GetCountByListRequest) (*pbmessage.GetCountByListResponse, error) {
	count, err := service.GetCountByList(request.ListId)
	if err != nil {
		logger.Error(err)
		return &pbmessage.GetCountByListResponse{}, err
	}

	return &pbmessage.GetCountByListResponse{Count: count.Message()}, nil
}

func (s *Server) GetDocumentsByListAndContexts(ctx context.Context, request *pbmessage.GetDocumentsByListAndContextsRequest) (*pbmessage.GetDocumentsByListAndContextsResponse, error) {
	documents, err := service.GetDocumentsByListAndContexts(request.ListId, request.ContextIds, request.Limit, request.Offset)
	if err != nil {
		logger.Error(err)
		return &pbmessage.GetDocumentsByListAndContextsResponse{}, err
	}

	return &pbmessage.GetDocumentsByListAndContextsResponse{Documents: documents.Message()}, nil
}

func (s *Server) GetCountByListAndContexts(ctx context.Context, request *pbmessage.GetCountByListAndContextsRequest) (*pbmessage.GetCountByListAndContextsResponse, error) {
	count, err := service.GetCountByListAndContexts(request.ListId, request.ContextIds)
	if err != nil {
		logger.Error(err)
		return &pbmessage.GetCountByListAndContextsResponse{}, err
	}

	return &pbmessage.GetCountByListAndContextsResponse{Count: count.Message()}, nil
}
