package common

import (
	"database/sql"
	"encoding/json"
)

type NullInt32 struct {
	sql.NullInt64
}

func (n *NullInt32) MarshalJSON() ([]byte, error) {
	if n.Valid == false {
		return []byte("null"), nil
	}
	return json.Marshal(n.Int32())
}

func (n *NullInt32) IsNull() bool {
	return n.Valid == false
}

func (n *NullInt32) Int32() int32 {
	if n.IsNull() {
		return 0
	} else {
		return int32(n.NullInt64.Int64)
	}
}
