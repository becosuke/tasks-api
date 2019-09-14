package task

import (
	"fmt"
	"strings"

	"github.com/becosuke/tasks-api/domain/entity/common"
	message "github.com/becosuke/tasks-api/protogen/message/task"
)

const Table = "task"
const Database = "tasks"
const PrimaryKey = "id"
const RelationalKey = "list_id"

const RecordCacheKey = "task_entity_%d"
const DocumentCacheKey = "task_document_%d"
const KeyAllCacheKey = "task_key_all_%d_%d" // limit, offset
const CountAllCacheKey = "task_count_all"
const KeyRelationalCacheKey = "task_key_relational_%d_%d_%d"                            // listId, limit, offset
const CountRelationalCacheKey = "task_count_relational_%d"                              // listId
const KeyRelationalAndPropertiesCacheKey = "task_key_relational_properties_%d_%s_%d_%d" // listId, contextIds, limit, offset
const CountRelationalAndPropertiesCacheKey = "task_count_relational_properties_%d_%s"   // listId, contextIds

func GetRecordCacheKey(id uint64) string {
	return fmt.Sprintf(RecordCacheKey, id)
}

func GetDocumentCacheKey(id uint64) string {
	return fmt.Sprintf(DocumentCacheKey, id)
}

func GetKeyAllCacheKey(limit uint32, offset uint32) string {
	return fmt.Sprintf(KeyAllCacheKey, limit, offset)
}

func GetCountAllCacheKey() string {
	return CountAllCacheKey
}

func GetKeyRelationalCacheKey(listId uint64, limit uint32, offset uint32) string {
	return fmt.Sprintf(KeyRelationalCacheKey, listId, limit, offset)
}

func GetCountRelationalCacheKey(listId uint64) string {
	return fmt.Sprintf(CountRelationalCacheKey, listId)
}

func GetKeyRelationalAndPropertiesCacheKey(listId uint64, contextIds []uint64, limit uint32, offset uint32) string {
	return fmt.Sprintf(
		KeyRelationalAndPropertiesCacheKey,
		listId,
		strings.Trim(strings.Replace(fmt.Sprint(contextIds), " ", ",", -1), "[]"),
		limit,
		offset,
	)
}

func GetCountRelationalAndPropertiesCacheKey(listId uint64, contextIds []uint64) string {
	return fmt.Sprintf(
		CountRelationalAndPropertiesCacheKey,
		listId,
		strings.Trim(strings.Replace(fmt.Sprint(contextIds), " ", ",", -1), "[]"),
	)
}

type Record struct {
	Id          uint64          `db:"id"`
	ListId      uint64          `db:"list_id"`
	Title       string          `db:"title"`
	CompletedAt common.Datetime `db:"completed_at"`
	CreatedAt   common.Datetime `db:"created_at"`
	UpdatedAt   common.Datetime `db:"updated_at"`
	DeletedAt   common.Datetime `db:"deleted_at"`
}

func (val *Record) Valid() bool {
	if val != nil && val.DeletedAt.IsNull() {
		return true
	}

	return false
}

func (val *Record) Completed() bool {
	if val != nil && !val.CompletedAt.IsNull() {
		return true
	}

	return false
}

type Document message.Document

func (val *Document) Valid() bool {
	if val != nil && val.Enabled == true {
		return true
	}

	return false
}

func (val *Document) Message() *message.Document {
	if val.Valid() {
		res := message.Document(*val)
		return &res
	}

	return &message.Document{}
}

type Documents []*Document

func (vals Documents) Message() []*message.Document {
	res := make([]*message.Document, 0, len(vals))
	for _, val := range vals {
		res = append(res, val.Message())
	}
	return res
}
