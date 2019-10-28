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

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	MysqlDialect = "MYSQL"
)

type MyStore struct {
	*sqlx.DB
}

type MyStoreConf struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DBName   string `json:"dbName" yaml:"dbName"`
}

func NewMyStore(conf MyStoreConf) *MyStore {

	lk := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.IP, conf.Port)

	lk += `?parseTime=true&loc=(Asia/Shanghai)&charset=utf8mb4`

	db, err := sqlx.Connect(MysqlDialect, lk)

	if err != nil {
		panic(errors.New("获取MYSQL连接异常"))
	}

	return &MyStore{DB: db}

}
