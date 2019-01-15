package context

import (
	"fmt"

	"github.com/becosuke/tasks-api/domain/entity/common"
	message "github.com/becosuke/tasks-api/protogen/message/context"
)

const Table = "context"
const Database = "tasks"
const PrimaryKey = "id"

const EntityCacheKey = "context_entity_%d"
const DocumentCacheKey = "context_document_%d"
const KeyAllCacheKey = "context_key_all_%d_%d"
const CountAllCacheKey = "context_count_all"

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
