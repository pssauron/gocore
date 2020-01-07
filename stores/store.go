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
	"strings"

	"github.com/pssauron/gocore/utils/strutils"

	"github.com/pssauron/gocore/libs"
)

const (
	dbTag      = "db"
	primaryTag = "primary"
)

type Table interface {
	TableName() string
}

type DBStoreConf struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DBName   string `json:"dbName" yaml:"dbName"`
	Idle     int    `json:"idle" yaml:"idle"`
	Active   int    `json:"active" yaml:"active"`
}

type RedisConf struct {
	Addr   string `json:"addr" yaml:"addr"`
	DB     int    `json:"db" yaml:"db"`
	Idle   int    `json:"idle" yaml:"idle"`
	Active int    `json:"active" yaml:"active"`
}

type field struct {
	*reflect.StructField
	IsPrimary bool
	IsEmpty   bool
	Col       string
	Value     interface{}
}

type sqlStruct struct {
	cols []string      //列名称
	args []interface{} //参数值
	pcol string        //主键列名
	pval interface{}   //主键列值
}

//获取Bean 中所有的Field
func getDBFields(bean interface{}) ([]field, error) {

	t := reflect.TypeOf(bean)

	v := reflect.Indirect(reflect.ValueOf(bean))

	return getFields(t, v)
}

func getFields(t reflect.Type, v reflect.Value) ([]field, error) {

	fs := make([]field, 0)

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

			f := field{}
			f.StructField = &sf
			f.Name = sf.Name
			f.Col = cn
			f.Value = v.FieldByName(sf.Name).Interface()
			f.IsPrimary = strings.ToUpper(f.Tag.Get(primaryTag)) == strings.ToUpper("true")
			f.IsEmpty = mpt
			fs = append(fs, f)
		}

	}

	return fs, nil
}

func getSqlStruct(fields []field) (*sqlStruct, error) {

	ss := sqlStruct{
		cols: make([]string, 0),
		args: make([]interface{}, 0),
		pcol: "",
	}

	for _, item := range fields {
		if item.IsPrimary {
			ss.pcol = item.Col
			ss.pval = item.Value
		} else if !item.IsEmpty {
			ss.cols = append(ss.cols, item.Col)
			ss.args = append(ss.args, item.Value)
		}

	}

	return &ss, nil
}

func getTableName(bean interface{}) string {
	if table, ok := bean.(Table); ok {
		return table.TableName()
	}
	t := reflect.TypeOf(bean)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return strutils.ToSnakeCase(t.Name())

}

func setValue(bean interface{}, value int64) error {
	fs, err := getDBFields(bean)
	if err != nil {
		return err
	}
	k, tp, err := getPrimary(fs)
	if err != nil {
		return err
	}
	v := reflect.Indirect(reflect.ValueOf(bean))

	switch tp.Type.String() {
	case "libs.Int":
		v.FieldByName(k).Set(reflect.ValueOf(libs.NewInt(int(value))))
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		v.FieldByName(k).SetInt(value)
	default:
		return errors.New("unsupported auto increment value")
	}
	return nil
}

func getPrimary(fs []field) (string, *reflect.StructField, error) {

	for _, item := range fs {
		if item.IsPrimary {
			return item.Name, item.StructField, nil
		}
	}

	return "", nil, errors.New("no primary key")

}
