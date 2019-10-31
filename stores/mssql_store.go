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

func NewMSStore(conf StoreConf) *MSStore {

	lk := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.IP, conf.Port, conf.DBName)

	lk += `?parseTime=true`

	db, err := sqlx.Connect(MysqlDialect, lk)

	if err != nil {
		panic(errors.New("获取SQL连接异常"))
	}

	return &MSStore{DB: db}

}
