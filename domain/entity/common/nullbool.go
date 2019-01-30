package common

import (
	"database/sql"
	"encoding/json"
)

type NullBool struct {
	sql.NullBool
}

func (n *NullBool) MarshalJSON() ([]byte, error) {
	if n.Valid == false {
		return []byte("null"), nil
	}

	return json.Marshal(n.Bool())
}

func (n *NullBool) IsNull() bool {
	return n.Valid == false
}

func (n *NullBool) Bool() bool {
	if n.IsNull() {
		return false
	}

	return n.NullBool.Bool
}
