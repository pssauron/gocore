//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/25 4:01 下午
//
//============================================================

package stores

import (
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

//MSStore sql server Store
type MSStore struct {
	*sqlx.DB
}

func NewMSStore(conf *DBStoreConf) *MSStore {

	u := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s;encrypt=disable", conf.IP, conf.Port, conf.User, conf.Password, conf.DBName)

	fmt.Println(u)

	db, err := sqlx.Open("sqlserver", u)

	if err != nil {
		panic(errors.New("获取SQL连接异常"))
	}

	return &MSStore{DB: db}
}

func GetSqlStruct(bean interface{}) (*sqlStruct, error) {
	fs, err := getDBFields(bean)
	if err != nil {
		return nil, err
	}

	return getSqlStruct(fs)
}
