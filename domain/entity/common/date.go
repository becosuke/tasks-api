package common

import (
	"encoding/json"
	"time"
)

const DateFormat = "2006-01-02"

type Date struct {
	Time
}

func (t Date) String() string {
	return t.Time.Format(DateFormat)
}

func (t Date) MarshalJSON() ([]byte, error) {
	if t.IsNull() {
		return []byte("null"), nil
	} else {
		return json.Marshal(t.String())
	}
}

func NewDate(t time.Time) Date {
	return Date{Time{t, true}}
}
