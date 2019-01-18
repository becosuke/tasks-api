package task

import (
	"github.com/becosuke/tasks-api/domain/entity/common"
	entity "github.com/becosuke/tasks-api/domain/entity/task"
	repository "github.com/becosuke/tasks-api/domain/repository/task"
)

func Create(listId uint64, title string) (*entity.Document, error) {
	return repository.Create(listId, title)
}

func Update(id uint64, listId uint64, title string) (*entity.Document, error) {
	return repository.Update(id, listId, title)
}

func Delete(id uint64) (*entity.Document, error) {
	return repository.Delete(id)
}

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

func GetDocumentsByList(listId uint64, limit int32, offset int32) ([]*entity.Document, error) {
	return repository.FetchDocumentsByRelationalKey(listId, limit, offset)
}

func GetCountByList(listId uint64) (*common.Count, error) {
	return repository.FetchCountByRelationalKey(listId)
}
