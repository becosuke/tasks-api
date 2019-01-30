package task

import (
	"fmt"

	"github.com/becosuke/tasks-api/domain/entity/common"
	message "github.com/becosuke/tasks-api/protogen/message/task"
)

const Table = "task"
const Database = "tasks"
const PrimaryKey = "id"
const RelationalKey = "list_id"

const EntityCacheKey = "task_entity_%d"
const DocumentCacheKey = "task_document_%d"
const KeyAllCacheKey = "task_key_all_%d_%d" // limit, offset
const CountAllCacheKey = "task_count_all"
const KeyRelationalCacheKey = "task_key_relational_%d_%d_%d" // listID, limit, offset
const CountRelationalCacheKey = "task_count_relational_%d"   // listID

func GetEntityCacheKey(id uint64) string {
	return fmt.Sprintf(EntityCacheKey, id)
}

func GetDocumentCacheKey(id uint64) string {
	return fmt.Sprintf(DocumentCacheKey, id)
}

func GetKeyAllCacheKey(limit int32, offset int32) string {
	return fmt.Sprintf(KeyAllCacheKey, limit, offset)
}

func GetCountAllCacheKey() string {
	return CountAllCacheKey
}

func GetKeyRelationalCacheKey(listID uint64, limit int32, offset int32) string {
	return fmt.Sprintf(KeyRelationalCacheKey, listID, limit, offset)
}

func GetCountRelationalCacheKey(listID uint64) string {
	return fmt.Sprintf(CountRelationalCacheKey, listID)
}

type Entity struct {
	ID          uint64          `db:"id"`
	ListID      uint64          `db:"list_id"`
	Title       string          `db:"title"`
	CreatedAt   common.Datetime `db:"created_at"`
	UpdatedAt   common.Datetime `db:"updated_at"`
	DeletedAt   common.Datetime `db:"deleted_at"`
	CompletedAt common.Datetime `db:"completed_at"`
}

func (val *Entity) Valid() bool {
	if val != nil && val.DeletedAt.IsNull() {
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
