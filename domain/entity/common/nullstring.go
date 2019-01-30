package common

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (n *NullString) MarshalJSON() ([]byte, error) {
	if n.Valid == false {
		return []byte("null"), nil
	}

	return json.Marshal(n.String())
}

func (n *NullString) IsNull() bool {
	return n.Valid == false
}

func (n *NullString) String() string {
	if n.IsNull() {
		return ""
	}

	return n.NullString.String
}
