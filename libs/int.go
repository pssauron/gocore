//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/24 5:20 下午
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

type Int struct {
	i sql.NullInt64
}

func NewInt(i int) Int {
	var ii Int
	ii.i.Valid = true
	ii.i.Int64 = int64(i)
	return ii
}

func (i *Int) Set(value int) {
	i.i.Int64 = int64(value)
	i.i.Valid = true
}

func (i Int) IsEmpty() bool {
	return i.i.Valid
}

func (i Int) Get() int {
	return int(i.i.Int64)
}

func (i Int) String() string {
	if !i.i.Valid {
		return "NULL"
	}
	return fmt.Sprint(i.i.Int64)
}

func (i Int) MarshalJSON() ([]byte, error) {
	if !i.i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.i.Int64)
}

func (i *Int) UnmarshalJSON(data []byte) error {
	var value *int64
	if err := json.Unmarshal(data, &value); err != nil {
		return nil
	}
	i.i.Valid = value != nil

	if value == nil {
		i.i.Int64 = 0
		return nil
	}
	i.i.Int64 = *value
	return nil
}

func (i Int) Value() (driver.Value, error) {
	if !i.i.Valid {
		return nil, nil
	}
	return i.i.Int64, nil
}

func (i *Int) Scan(data interface{}) error {
	return i.i.Scan(data)
}
