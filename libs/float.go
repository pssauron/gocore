//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/24 5:05 下午
//
//
//============================================================

package libs

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Float struct {
	f sql.NullFloat64
}

func NewFloat(value float64) Float {
	var f Float
	f.f.Valid = true
	f.f.Float64 = value
	return f
}

func (f Float) IsEmpty() bool {
	return f.f.Valid
}

func (f Float) Get() float64 {
	return f.f.Float64
}

func (f Float) String() string {
	if !f.f.Valid {
		return "NULL"
	}
	return fmt.Sprint(f.f.Float64)
}

func (f Float) Value() (driver.Value, error) {
	if !f.f.Valid {
		return nil, nil
	}
	return f.f.Float64, nil
}

func (f *Float) Scan(data interface{}) error {
	return f.f.Scan(data)
}

func (f Float) MarshalJSON() ([]byte, error) {
	if !f.f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.f.Float64)

}

func (f *Float) UnmarshalJSON(data []byte) error {

	var value *float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	f.f.Valid = value != nil

	if value == nil {
		f.f.Float64 = 0
		return nil
	}
	f.f.Float64 = *value
	return nil

}
