package task

import (
	entity "github.com/becosuke/tasks-api/domain/entity/task"
)

func Document(val *entity.Record) *entity.Document {
	if val.Valid() == false {
		return &entity.Document{}
	}

	res := &entity.Document{
		Id:        val.Id,
		ListId:    val.ListId,
		Title:     val.Title,
		Completed: val.Completed(),
		Enabled:   val.Valid(),
		CreatedAt: val.CreatedAt.Unix(),
		UpdatedAt: val.UpdatedAt.Unix(),
		DeletedAt: val.DeletedAt.Unix(),
	}

	return res
}
