package task

import (
	entity "github.com/becosuke/tasks-api/domain/entity/task"
)

func Document(val *entity.Entity) *entity.Document {
	if val.Valid() == false {
		return &entity.Document{}
	}

	res := &entity.Document{
		Id:        val.ID,
		ListId:    val.ListID,
		Title:     val.Title,
		Enabled:   val.Valid(),
		CreatedAt: val.CreatedAt.Unix(),
		UpdatedAt: val.UpdatedAt.Unix(),
		DeletedAt: val.DeletedAt.Unix(),
	}

	return res
}
