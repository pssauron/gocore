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

	"github.com/jmoiron/sqlx"
)

//MSStore sql server Store
type MSStore struct {
	db *sqlx.DB
}

func NewMSStore(conf StoreConf) *MSStore {

	lk := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.IP, conf.Port, conf.DBName)

	lk += `?parseTime=true`

	db, err := sqlx.Connect(MysqlDialect, lk)

	if err != nil {
		panic(errors.New("获取MYSQL连接异常"))
	}

	return &MSStore{db: db}

}
