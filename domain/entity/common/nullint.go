package common

import (
	"database/sql"
	"encoding/json"
	"strconv"
)

type NullInt struct {
	sql.NullInt64
}

func (n *NullInt) MarshalJSON() ([]byte, error) {
	if n.Valid == false {
		return []byte("null"), nil
	}
	return json.Marshal(n.Int())
}

func (n *NullInt) IsNull() bool {
	return n.Valid == false
}

func (n *NullInt) Int() int {
	if n.IsNull() {
		return 0
	} else {
		return int(n.NullInt64.Int64)
	}
}

func (n *NullInt) String() string {
	if n.IsNull() {
		return ""
	} else {
		return strconv.FormatInt(n.NullInt64.Int64, 10)
	}
}
