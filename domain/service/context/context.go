package context

import (
	"github.com/becosuke/tasks-api/domain/entity/common"
	entity "github.com/becosuke/tasks-api/domain/entity/context"
	repository "github.com/becosuke/tasks-api/domain/repository/context"
)

func GetDocument(id uint64) (*entity.Document, error) {
	return repository.FetchDocument(id)
}

func GetDocuments(ids []uint64) ([]*entity.Document, error) {
	return repository.FetchDocuments(ids)
}

func GetDocumentsAll(limit int32, offset int32) ([]*entity.Document, error) {
	return repository.FetchDocumentsAll(limit, offset)
}

func GetCountAll() (*common.Count, error) {
	return repository.FetchCountAll()
}
