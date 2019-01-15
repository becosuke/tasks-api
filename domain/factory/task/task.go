package task

import (
	entity "github.com/becosuke/tasks-api/domain/entity/task"
)

func Document(val *entity.Entity) *entity.Document {
	if val.Valid() == false {
		return &entity.Document{}
	}

	var enabled bool
	if val.DeletedAt.IsNull() {
		enabled = true
	} else {
		enabled = false
	}

	res := &entity.Document{
		Id:      val.ID,
		Title:   val.Title,
		Enabled: enabled,
	}

	return res
}
