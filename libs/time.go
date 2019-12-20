//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/24 5:43 下午
//
//
//============================================================
package libs

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type Time struct {
	t sql.NullTime
}

func NewTime(time time.Time) Time {
	var t Time
	t.t.Valid = true
	t.t.Time = time
	return t
}

func (t Time) IsEmpty() bool {
	return !t.t.Valid || t.t.Time == time.Unix(0, 0)
}

func (t Time) Get() time.Time {
	return t.t.Time
}

func (t Time) String() string {
	if !t.t.Valid {
		return "NULL"
	}
	return t.t.Time.Format(timeFormat)
}

func (t *Time) Scan(value interface{}) error {
	return t.t.Scan(value)
}

func (t Time) Value() (driver.Value, error) {
	if !t.t.Valid {
		return nil, nil
	}
	return t.t.Time, nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	if !t.t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.t.Time.Format(timeFormat))
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var value interface{}

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	switch value.(type) {
	case string:
		return t.UnmarshalText([]byte(value.(string)))
	case nil:
		t.t.Time = time.Time{}
		t.t.Valid = false
		return nil
	default:
		return errors.New("不支持的序列化类型")
	}
}

//UnmarshalText json 反序列化
func (t *Time) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		t.t.Valid = false
		return nil
	}
	tt, err := time.Parse(timeFormat, str)

	if err != nil {
		return err
	}
	t.t.Time = tt
	t.t.Valid = true
	return nil
}
