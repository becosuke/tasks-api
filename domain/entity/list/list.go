package list

import (
	"fmt"

	"github.com/becosuke/tasks-api/domain/entity/common"
	message "github.com/becosuke/tasks-api/protogen/message/list"
)

const Table = "list"
const Database = "tasks"
const PrimaryKey = "id"

const EntityCacheKey = "list_entity_%d"
const DocumentCacheKey = "list_document_%d"
const KeyAllCacheKey = "list_key_all_%d_%d"
const CountAllCacheKey = "list_count_all"

func GetEntityCacheKey(id uint64) string {
	return fmt.Sprintf(EntityCacheKey, id)
}

func GetDocumentCacheKey(id uint64) string {
	return fmt.Sprintf(DocumentCacheKey, id)
}

func GetKeyAllCacheKey(limit int32, offset int32) string {
	return KeyAllCacheKey
}

func GetCountAllCacheKey() string {
	return CountAllCacheKey
}

type Entity struct {
	ID        uint64          `db:"id"`
	Title     string          `db:"title"`
	CreatedAt common.Datetime `db:"created_at"`
	UpdatedAt common.Datetime `db:"updated_at"`
	DeletedAt common.Datetime `db:"deleted_at"`
}

func (val *Entity) Valid() bool {
	if val != nil && val.DeletedAt.IsNull() {
		return true
	} else {
		return false
	}
}

type Document message.Document

func (val *Document) Valid() bool {
	if val != nil && val.Enabled == true {
		return true
	} else {
		return false
	}
}

func (val *Document) Message() *message.Document {
	if val.Valid() {
		res := message.Document(*val)
		return &res
	} else {
		return &message.Document{}
	}
}
