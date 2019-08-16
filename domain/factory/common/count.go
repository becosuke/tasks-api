package common

import (
	entity "github.com/becosuke/tasks-api/domain/entity/common"
)

func Count(count uint64) *entity.Count {
	return &entity.Count{
		Count:   count,
		Enabled: true,
	}
}
