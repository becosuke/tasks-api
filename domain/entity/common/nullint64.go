package common

import (
	"database/sql"
	"encoding/json"
)

type NullInt64 struct {
	sql.NullInt64
}

func (n *NullInt64) MarshalJSON() ([]byte, error) {
	if n.Valid == false {
		return []byte("null"), nil
	}
	return json.Marshal(n.Int64())
}

func (n *NullInt64) IsNull() bool {
	return n.Valid == false
}

func (n *NullInt64) Int64() int64 {
	if n.IsNull() {
		return 0
	} else {
		return n.NullInt64.Int64
	}
}
