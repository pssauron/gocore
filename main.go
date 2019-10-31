//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/30 5:02 下午
//
//============================================================

package main

import (
	"fmt"

	"github.com/pssauron/gocore/utils/strutils"

	"github.com/pssauron/gocore/libs"
)

type SysLogin struct {
	LoginID   libs.Int    `db:"LoginID" json:"loginId" primary:"true"`
	UserAcct  libs.String `db:"UserAcct" json:"userAcct"`
	Mobile    libs.String `db:"Mobile" json:"mobile"`
	Password  libs.String `db:"Password" json:"password"`
	RegDate   libs.Time   `db:"RegDate" json:"regDate"`
	LoginIP   libs.String `db:"LoginIP" json:"loginIp"`
	LoginTime libs.Time   `db:"LoginTime" json:"loginTime"`
	DR        libs.Bool   `db:"DR" json:"dr"`
}

func (SysLogin) TableName() string {
	return "SysLogin"
}

func main() {

	//conf := stores.MyStoreConf{
	//	IP:       "192.168.0.110",
	//	Port:     "3306",
	//	User:     "root",
	//	Password: "123456",
	//	DBName:   "yunlian-platform",
	//}

	//db := stores.NewMyStore(conf)
	//login := &SysLogin{
	//	LoginID:  libs.NewInt(100058),
	//	UserAcct: libs.NewString("aaaaaassss"),
	//	Mobile:   libs.NewString("18682331124"),
	//	Password: libs.NewString("222222"),
	//}
	//err := db.Update(login)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//list := make([]SysLogin, 0)

	//err := db.Query(&list, "select * from SysLogin")
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(len(list))

	//pagedata, err := db.QueryPage(&list, "select * from SysLogin", 1, 2)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//bs, err := json.Marshal(pagedata)
	//
	//fmt.Println(string(bs))

	fmt.Println(strutils.ToSnakeCase("SysLogin"))

}
