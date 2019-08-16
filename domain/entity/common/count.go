package common

import (
	pbmessage "github.com/becosuke/tasks-api/protogen/message/common"
)

type Count pbmessage.Count

func (val *Count) Valid() bool {
	return val != nil
}

func (val *Count) Message() *pbmessage.Count {
	if !val.Valid() {
		return &pbmessage.Count{}
	}
	return (*pbmessage.Count)(val)
}
