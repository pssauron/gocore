//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/25 12:04 下午
//
//============================================================
package stores

import (
	"reflect"
	"strings"
)

type (
	//table Name 表示 对象对应 table
	Storable interface {
		TableName() string
	}
)

const (
	dbTag      = "db"
	autoTag    = "auto"
	primaryTag = "primary"
)

const (
	stringName = "string"
	boolName   = "bool"
)

type StorableMeta struct {
	ID        *Field
	Fields    []Field
	TableName string
}

type Field struct {
	Name      string      //字段Struct filed name
	Tag       string      //字段 sql name
	Value     interface{} //字段 value
	IsPrimary bool        //是否主键
	IsAuto    bool        //是否自增
}

func getMeta(s Storable) (*StorableMeta, error) {

	m := new(StorableMeta)

	m.TableName = s.TableName()

	t := reflect.TypeOf(s)

	v := reflect.Indirect(reflect.ValueOf(s))

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		dt := f.Tag.Get(dbTag)
		if dt == "-" {
			continue
		}

		getFields(f, v, m)

	}

	return m, nil
}

func getFields(f reflect.StructField, v reflect.Value, m *StorableMeta) error {
	dt := f.Tag.Get(dbTag)
	if dt == "" {
		t := f.Type
		for i := 0; i < t.NumField(); i++ {
			ff := t.Field(i)
			if ff.Tag.Get(dbTag) == "-" {
				continue
			}
			err := getFields(ff, v, m)
			if err != nil {
				return err
			}
		}

	}

	field := Field{}
	field.Name = f.Name
	field.Value = v.FieldByName(f.Name).Interface()
	field.Tag = dt
	if f.Tag.Get(primaryTag) != "" && "TRUE" == strings.ToUpper(f.Tag.Get(primaryTag)) {

		field.IsPrimary = true

		if f.Tag.Get(autoTag) != "" && "TRUE" == strings.ToUpper(f.Tag.Get(autoTag)) {
			field.IsAuto = true
		}

	}

	if field.IsPrimary {
		m.ID = &field
	} else {
		m.Fields = append(m.Fields, field)
	}

	return nil

}
