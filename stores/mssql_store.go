//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/25 4:01 下午
//
//============================================================

package stores

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/pssauron/gocore/libs"
	"github.com/pssauron/gocore/rs"
	"github.com/pssauron/log4go"

	"github.com/jmoiron/sqlx"
)

//MSStore sql server Store
type MSStore struct {
	*sqlx.DB
}

type PageData struct {
}

func NewMSStore(conf *DBStoreConf) *MSStore {

	u := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s;encrypt=disable", conf.IP, conf.Port, conf.User, conf.Password, conf.DBName)

	fmt.Println(u)

	db, err := sqlx.Open("mssql", u)

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

//Insert 带事务保存
func (ms *MSStore) Insert(tx *sql.Tx, bean interface{}) error {

	err := ms.insert(tx, bean)

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil

}

//Insert 批量保存
func (ms *MSStore) Insertx(tx *sql.Tx, beans ...interface{}) error {

	for _, item := range beans {
		err := ms.insert(tx, item)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil

}

//insert 保存
func (ms *MSStore) insert(tx *sql.Tx, bean interface{}) error {

	//
	st, err := GetSqlStruct(bean)

	if err != nil {
		return err
	}

	cols := st.cols
	args := st.args

	args = append(args, st.pval)
	cols = append(cols, st.pcol)

	cs := strings.Join(cols, ",")

	phs := make([]string, 0)

	for _, _ = range args {
		phs = append(phs, "?")
	}

	q := `INSERT INTO ` + getTableName(bean) + "(" + cs + ") VALUES (" + strings.Join(phs, ",") + ")"

	log4go.Debug("执行SQL:%s", q)

	_, err = tx.Exec(q, args...)

	if err != nil {
		return err
	}

	return nil
}

func (ms *MSStore) Update(tx *sql.Tx, bean interface{}) error {
	err := ms.update(tx, bean)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (ms *MSStore) Updatex(tx *sql.Tx, beans ...interface{}) error {
	for _, item := range beans {
		err := ms.update(tx, item)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}

// update 修改语句
func (ms *MSStore) update(tx *sql.Tx, bean interface{}) error {

	//
	st, err := GetSqlStruct(bean)

	if err != nil {
		return err
	}

	q := "UPDATE " + getTableName(bean) + "  SET "

	for idx, item := range st.cols {
		q += item + " = ?"
		if idx != len(st.cols)-1 {
			q += ", "
		}
	}

	q += " WHERE a." + st.pcol + " = ?"

	args := st.args

	args = append(args, st.pval)

	log4go.Debug("【执行更新语句:%s 】", q)

	_, err = tx.Exec(q, args...)

	if err != nil {
		return err
	}

	return nil
}

func (ms *MSStore) QueryPageWithOffset(dest interface{}, query string, page, size int, args ...interface{}) (*rs.PageData, error) {

	q := `select count(*) from (` + query + ` a`

	var count int

	err := ms.Get(&count, q, args...)

	if err != nil {
		return nil, err
	}

	pdata := rs.PageData{}

	pdata.Totals = libs.NewInt(count)
	pdata.Page = libs.NewInt(page)
	pdata.Size = libs.NewInt(size)
	if count%size == 0 {
		pdata.Pages = libs.NewInt(count / size)
	} else {
		pdata.Pages = libs.NewInt(count/size + 1)
	}
	if page > pdata.Pages.Get() {
		return nil, errors.New("分页超过最大页数")
	}

	q = `select * from (` + query + `) a offset ? rows fetch next ? rows only`

	args = append(args, (page-1)*size, size)

	err = ms.Select(dest, q, args...)

	if err != nil {
		return nil, err
	}

	return &pdata, nil
}
