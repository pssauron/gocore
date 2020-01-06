//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/30 5:02 下午
//
//============================================================

package main

import (
	"fmt"

	"github.com/pssauron/gocore/libs"
	"github.com/pssauron/gocore/utils/strutils"
)

type SEOrder struct {
	FBrNo    libs.String `db:"FBrNo"`
	FInterID libs.Int    `db:"FInterID"`
	FBillNo  libs.String `db:"FBillNo"`
}

func (SEOrder) TableName() string {
	return "SEOrder"
}

func main() {

	key := "thikeives"

	str1, err := strutils.AesEncrypt("sa@123", key)

	if err != nil {
		panic(err)
	}

	fmt.Println(str1)

	str, err := strutils.AesDecrypt(str1, key)
	if err != nil {
		panic(err)
	}

	fmt.Println(str)

	//conf := stores.DBStoreConf{
	//	IP:       "47.107.101.69",
	//	Port:     "1433",
	//	User:     "sa",
	//	Password: "sa@123",
	//	DBName:   "AIS201407281401391",
	//}
	//
	//db := stores.NewMSStore(&conf)
	//orders := make([]SEOrder, 0)
	//err := db.Select(&orders, "select  * from SEOrder offset 10 rows ")
	//if err != nil {
	//	panic(err)
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

}
