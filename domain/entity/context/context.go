package context

import (
	"fmt"

	"github.com/becosuke/tasks-api/domain/entity/common"
	pbmessage "github.com/becosuke/tasks-api/protogen/message/context"
)

const Table = "context"
const Database = "tasks"
const PrimaryKey = "id"

const RecordCacheKey = "context_record_%d"
const KeyAllCacheKey = "context_key_all_%d_%d" // limit, offset
const CountAllCacheKey = "context_count_all"

func GetRecordCacheKey(id uint64) string {
	return fmt.Sprintf(RecordCacheKey, id)
}

func GetKeyAllCacheKey(limit uint32, offset uint32) string {
	return KeyAllCacheKey
}

func GetCountAllCacheKey() string {
	return CountAllCacheKey
}

type Record struct {
	Id        uint64          `db:"id"`
	Title     string          `db:"title"`
	CreatedAt common.Datetime `db:"created_at"`
	UpdatedAt common.Datetime `db:"updated_at"`
	DeletedAt common.Datetime `db:"deleted_at"`
}

func (val *Record) Valid() bool {
	return val != nil && val.DeletedAt.IsNull()
}

type Document pbmessage.Document

func (val *Document) Valid() bool {
	return val != nil && val.Enabled
}

func (val *Document) Message() *pbmessage.Document {
	if !val.Valid() {
		return &pbmessage.Document{}
	}

	return (*pbmessage.Document)(val)
}

type Documents []*Document

func (vals *Documents) Valid() bool {
	return vals != nil && len(*vals) > 0
}

func (vals *Documents) Message() []*pbmessage.Document {
	if !vals.Valid() {
		return make([]*pbmessage.Document, 0)
	}

	res := make([]*pbmessage.Document, 0, len(*vals))
	for _, val := range *vals {
		if !val.Valid() {
			continue
		}
		res = append(res, val.Message())
	}
	return res
}
