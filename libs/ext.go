//============================================================
// 描述: 数据库json 类型
// 作者: Simon
// 日期: 2019/10/30 4:41 下午
//
//============================================================

package libs

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Ext struct {
	m map[string]string
}

func NewExt() Ext {
	return Ext{m: map[string]string{}}
}

func (e *Ext) SetExtItem(key string, value string) {
	e.m[key] = value
}

func (e Ext) Get() map[string]string {
	return e.m
}

func (e Ext) IsEmpty() bool {
	return e.m == nil || len(e.m) == 0
}

func (e *Ext) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		err := json.Unmarshal(v, &e)
		return err
	case string:
		return json.Unmarshal([]byte(v), &e)
	default:
		return errors.New("scan 拓展字段 Ext 异常")
	}

	return nil
}

func (e Ext) Value() (driver.Value, error) {
	return json.Marshal(e)
}

func (e Ext) MarshalJSON() ([]byte, error) {

	return json.Marshal(e.m)
}

func (e *Ext) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e.m)
}
