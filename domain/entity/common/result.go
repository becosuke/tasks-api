package common

import (
	message "github.com/becosuke/tasks-api/protogen/message/common"
)

type Result message.Result

func (val *Result) Valid() bool {
	if val != nil {
		return true
	} else {
		return false
	}
}

func (val *Result) Message() *message.Result {
	if val.Valid() {
		res := message.Result(*val)
		return &res
	} else {
		return &message.Result{}
	}
}
