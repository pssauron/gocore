//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/24 5:34 下午
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

type Bool struct {
	b sql.NullBool
}

func NewBool(value bool) Bool {
	var b Bool
	b.b.Valid = true
	b.b.Bool = value
	return b
}

func (b Bool) IsEmpty() bool {
	return !b.b.Valid
}

func (b Bool) Get() bool {
	return b.b.Bool
}

func (b Bool) String() string {
	if !b.b.Valid {
		return "NULL"
	}
	return fmt.Sprint(b.b.Bool)
}

func (b Bool) Value() (driver.Value, error) {
	if !b.b.Valid {
		return nil, nil
	}
	return b.b.Bool, nil
}

func (b *Bool) Scan(data interface{}) error {
	return b.b.Scan(data)
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.b.Bool)
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	var bb *bool
	if err := json.Unmarshal(data, &bb); err != nil {
		return err
	}
	b.b.Valid = bb != nil
	if bb == nil {
		b.b.Bool = false
		return nil
	}
	b.b.Bool = *bb
	return nil
}
