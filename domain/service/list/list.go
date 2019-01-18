package list

import (
	entity "github.com/becosuke/tasks-api/domain/entity/list"
	repository "github.com/becosuke/tasks-api/domain/repository/list"
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

func Create(title string) (*entity.Document, error) {
	return repository.Create(title)
}

func Update(id uint64, title string) (*entity.Document, error) {
	return repository.Update(id, title)
}

func Delete(id uint64) (*entity.Document, error) {
	return repository.Delete(id)
}
