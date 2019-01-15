package common

import (
	"database/sql"
	"encoding/json"
)

type NullUint struct {
	sql.NullInt64
}

func (n *NullUint) MarshalJSON() ([]byte, error) {
	if n.Valid == false {
		return []byte("null"), nil
	}
	return json.Marshal(n.Uint())
}

func (n *NullUint) IsNull() bool {
	return n.Valid == false
}

func (n *NullUint) Uint() uint {
	if n.IsNull() {
		return 0
	} else {
		return uint(n.NullInt64.Int64)
	}
}
