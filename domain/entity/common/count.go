package common

import (
	message "github.com/becosuke/tasks-api/protogen/message/common"
)

type Count message.Count

func (val *Count) Valid() bool {
	if val != nil {
		return true
	}

	return false
}

func (val *Count) Message() *message.Count {
	if val.Valid() {
		res := message.Count(*val)
		return &res
	}

	return &message.Count{}
}

func NewCount(count uint64) *Count {
	return &Count{Count: count}
}
