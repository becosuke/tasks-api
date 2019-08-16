package list

import (
	entity "github.com/becosuke/tasks-api/domain/entity/list"
)

func Document(val *entity.Record) *entity.Document {
	if !val.Valid() {
		return &entity.Document{}
	}

	res := &entity.Document{
		Id:        val.Id,
		Title:     val.Title,
		Enabled:   true,
		CreatedAt: val.CreatedAt.Unix(),
		UpdatedAt: val.UpdatedAt.Unix(),
		DeletedAt: val.DeletedAt.Unix(),
	}

	return res
}
