package context

import (
	"github.com/becosuke/tasks-api/domain/entity/common"
	entity "github.com/becosuke/tasks-api/domain/entity/context"
	repository "github.com/becosuke/tasks-api/domain/repository/context"
)

func Create(title string) (*entity.Document, error) {
	document, err := repository.Create(title)
	if err != nil {
		return &entity.Document{}, err
	}
	return document, nil
}

func Update(id uint64, title string) (*entity.Document, error) {
	document, err := repository.Update(id, title)
	if err != nil {
		return &entity.Document{}, err
	}
	return document, nil
}

func Delete(id uint64) (*entity.Document, error) {
	document, err := repository.Delete(id)
	if err != nil {
		return &entity.Document{}, err
	}
	return document, nil
}

func GetDocument(id uint64) (*entity.Document, error) {
	document, err := repository.FetchDocument(id)
	if err != nil {
		return &entity.Document{}, err
	}
	return document, nil
}

func GetDocuments(ids []uint64) (entity.Documents, error) {
	documents, err := repository.FetchDocuments(ids)
	if err != nil {
		return entity.Documents{}, nil
	}
	return documents, nil
}

func GetDocumentsAll(limit uint32, offset uint32) (entity.Documents, error) {
	documents, err := repository.FetchDocumentsAll(limit, offset)
	if err != nil {
		return entity.Documents{}, nil
	}
	return documents, nil
}

func GetCountAll() (*common.Count, error) {
	count, err := repository.FetchCountAll()
	if err != nil {
		return &common.Count{}, nil
	}
	return count, nil
}
