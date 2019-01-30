package common

import (
	"encoding/json"
	"time"
)

const DatetimeFormat = "2006-01-02 15:04:05"

type Datetime struct {
	Time
}

func (t Datetime) String() string {
	return t.Time.Format(DatetimeFormat)
}

func (t Datetime) MarshalJSON() ([]byte, error) {
	if t.IsNull() {
		return []byte("null"), nil
	}

	return json.Marshal(t.String())
}

func NewDatetime(t time.Time) Datetime {
	return Datetime{Time{t, true}}
}
