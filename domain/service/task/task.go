package task

import (
	entity "github.com/becosuke/tasks-api/domain/entity/task"
	repository "github.com/becosuke/tasks-api/domain/repository/task"
)

func GetDocument(id uint64) (*entity.Document, error) {
	return repository.FetchDocument(id)
}

func GetDocuments(ids []uint64) ([]*entity.Document, error) {
	return repository.FetchDocuments(ids)
}

func GetDocumentsAll(limit int32, offset int32) ([]*entity.Document, error) {
	return repository.FetchAll(limit, offset)
}

func GetCountAll() (uint64, error) {
	return repository.CountAll()
}