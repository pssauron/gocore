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
)

type String struct {
	s sql.NullString
}

func NewString(value string) String {
	var s String
	s.s.Valid = true
	s.s.String = value
	return s
}
func (s *String) Scan(value interface{}) error {
	return s.s.Scan(value)
}
func (s String) Value() (driver.Value, error) {
	if !s.s.Valid {
		return nil, nil
	}
	return s.s.String, nil
}
func (s *String) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	s.s.Valid = value != nil

	if value == nil {
		s.s.String = ""
		return nil
	}

	s.s.String = *value

	return nil
}

func (s String) MarshalJSON() ([]byte, error) {
	if !s.s.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.s.String)
}

func (s String) IsNil() bool {
	return !s.s.Valid || s.s.String == ""
}

func (s String) Get() string {
	return s.s.String
}

func (s String) String() string {
	if !s.s.Valid {
		return "NULL"
	}
	return s.s.String
}
