package common

import (
	"database/sql"
	"encoding/json"
)

type NullFloat64 struct {
	sql.NullFloat64
}

func (n *NullFloat64) MarshalJSON() ([]byte, error) {
	if n.Valid == false {
		return []byte("null"), nil
	}

	return json.Marshal(n.Float64())
}

func (n *NullFloat64) IsNull() bool {
	return n.Valid == false
}

func (n *NullFloat64) Float64() float64 {
	if n.IsNull() {
		return 0.0
	}

	return n.NullFloat64.Float64
}
