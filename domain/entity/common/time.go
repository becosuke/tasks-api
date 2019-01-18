package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/becosuke/tasks-api/config"
)

var EpochTime = time.Unix(0, 0)
var EpochTimestamp int64 = 0

type Time struct {
	time.Time
	Valid bool
}

func (t *Time) Scan(value interface{}) (err error) {
	var s mysql.NullTime
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*t = Time{EpochTime, false}
	} else {
		*t = Time{s.Time, true}
	}

	return nil
}

func (t Time) Unix() int64 {
	if t.IsNull() {
		return EpochTimestamp
	} else {
		return t.Time.Unix()
	}
}

func (t Time) Value() (driver.Value, error) {
	if t.IsNull() {
		return nil, nil
	} else {
		return t.Time.String(), nil
	}
}

func (t Time) IsNull() bool {
	return t.Valid == false
}

func (t Time) ElaspedTime() string {
	if t.IsNull() {
		return ""
	}

	var elasped int64
	conf := config.GetConfig()
	now := conf.NowTimestamp
	if elasped = now - t.Time.Unix(); elasped < 0 {
		elasped = 0
	}

	if elasped < 86400 {
		var hour int64 = 3600
		var min int64 = 60
		if elasped == 0 {
			return "たった今"
		} else if elasped < min {
			return fmt.Sprintf("%d秒前", elasped)
		} else if elasped < hour {
			return fmt.Sprintf("%d分前", elasped/min)
		} else {
			return fmt.Sprintf("%d時間前", elasped/hour)
		}
	} else {
		return t.Time.Format("2006年1月2日")
	}
}

func (t Time) String() string {
	return t.Time.Format(time.RFC3339)
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsNull() {
		return []byte("null"), nil
	} else {
		return json.Marshal(t.Time.Unix())
	}
}
