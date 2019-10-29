//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/25 12:04 下午
//
//============================================================
package stores

import (
	"errors"
	"reflect"

	"github.com/pssauron/gocore/libs"
)

const (
	dbTag      = "db"
	primaryTag = "primary"
)

type Table interface {
	TableName() string
}

type Field struct {
	*reflect.StructField
	IsPrimary bool
	IsEmpty   bool
	Col       string
	Value     interface{}
}

//获取Bean 中所有的Field
func GetDBFields(bean interface{}) ([]Field, error) {

	t := reflect.TypeOf(bean)

	v := reflect.Indirect(reflect.ValueOf(bean))

	return getFields(t, v)
}

func getFields(t reflect.Type, v reflect.Value) ([]Field, error) {

	fs := make([]Field, 0)

	if t.Kind() == reflect.Ptr {
		return getFields(t.Elem(), v)
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("不支持非结构体类型")
	}

	cnt := t.NumField()

	for i := 0; i < cnt; i++ {
		sf := t.Field(i)
		ft := sf.Type
		fv := v.Field(i)
		cn := sf.Tag.Get(dbTag)
		//如果 tag 为 "-" 忽略
		if cn == "-" {
			continue
		}
		if cn == "" {
			cfs, err := getFields(ft, reflect.Indirect(v.FieldByName(sf.Name)))

			if err != nil {
				return nil, err
			}

			fs = append(fs, cfs...)
		} else {
			var mpt bool

			switch sf.Type.String() {
			case "libs.String", "libs.Int", "libs.Bool", "libs.Float", "libs.Time", "libs.Ext":

				mpt = fv.Interface().(libs.LibItem).IsEmpty()

			default:
				mpt = fv.IsZero()
			}

			f := Field{}
			f.StructField = &sf
			f.Name = sf.Name
			f.Col = cn
			f.Value = v.FieldByName(sf.Name).Interface()
			f.IsPrimary = f.Tag.Get(primaryTag) == "true"
			f.IsEmpty = mpt
			fs = append(fs, f)
		}

	}

	return fs, nil
}
