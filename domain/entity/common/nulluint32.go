package common

import (
	"database/sql"
	"encoding/json"
)

type NullUint32 struct {
	sql.NullInt64
}

func (n *NullUint32) MarshalJSON() ([]byte, error) {
	if n.Valid == false {
		return []byte("null"), nil
	}
	return json.Marshal(n.Uint32())
}

func (n *NullUint32) IsNull() bool {
	return n.Valid == false
}

func (n *NullUint32) Uint32() uint32 {
	if n.IsNull() {
		return 0
	} else {
		return uint32(n.NullInt64.Int64)
	}
}
