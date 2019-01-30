package common

import (
	"database/sql"
	"encoding/json"
)

type NullUint64 struct {
	sql.NullInt64
}

func (n *NullUint64) MarshalJSON() ([]byte, error) {
	if n.Valid == false {
		return []byte("null"), nil
	}

	return json.Marshal(n.Uint64())
}

func (n *NullUint64) IsNull() bool {
	return n.Valid == false
}

func (n *NullUint64) Uint64() uint64 {
	if n.IsNull() {
		return 0
	}

	return uint64(n.NullInt64.Int64)
}
