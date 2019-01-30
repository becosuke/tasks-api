package common

import (
	"database/sql"
	"encoding/json"
)

type NullFloat32 struct {
	sql.NullFloat64
}

func (n *NullFloat32) MarshalJSON() ([]byte, error) {
	if n.Valid == false {
		return []byte("null"), nil
	}

	return json.Marshal(n.Float32())
}

func (n *NullFloat32) IsNull() bool {
	return n.Valid == false
}

func (n *NullFloat32) Float32() float32 {
	if n.IsNull() {
		return 0.0
	}

	return float32(n.NullFloat64.Float64)
}
