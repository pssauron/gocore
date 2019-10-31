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

	"github.com/pssauron/gocore/libs"

	"github.com/pssauron/gocore/rs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	MysqlDialect = "mysql"
)

type MyStore struct {
	db *sqlx.DB
}

func NewMyStore(conf StoreConf) *MyStore {

	lk := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.IP, conf.Port, conf.DBName)

	lk += `?parseTime=true`

	db, err := sqlx.Connect(MysqlDialect, lk)

	if err != nil {
		panic(errors.New("获取MYSQL连接异常"))
	}

	return &MyStore{db: db}

}

func (m *MyStore) BeginTx() (*sql.Tx, error) {
	return m.db.Begin()
}

func (m *MyStore) RollBack(tx *sql.Tx) error {
	return tx.Rollback()
}

func (m *MyStore) Commit(tx *sql.Tx) error {
	return tx.Rollback()
}

func (m *MyStore) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.db.Exec(query, args...)
}

func (m *MyStore) Query(dest interface{}, sql string, args ...interface{}) error {
	return m.db.Select(dest, sql, args...)
}

func (m *MyStore) Get(dest interface{}, sql string, args ...interface{}) error {
	return m.db.Get(dest, sql, args...)
}

func (m *MyStore) QueryPage(dest interface{}, query string, page, size int, args ...interface{}) (*rs.PageData, error) {

	//查询分页条数
	var count int

	cs := `select count(*) from (%s) a`

	err := m.db.Get(&count, fmt.Sprintf(cs, query), args...)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return rs.NewPageData(page, size), nil
	}

	ds := `select a.* from (%s) a limit ?,?`

	args = append(args, (page-1)*size, size)

	err = m.db.Select(dest, fmt.Sprintf(ds, query), args...)

	if err != nil {
		return nil, err
	}

	po := rs.NewPageData(page, size)

	po.Totals = libs.NewInt(count)

	if count%size == 0 {
		po.Pages = libs.NewInt(count / size)
	} else {
		po.Pages = libs.NewInt(count/size + 1)
	}

	po.Data = dest

	return po, nil
}

func (m *MyStore) Insert(bean interface{}) error {
	tx, err := m.BeginTx()
	if err != nil {
		return err
	}
	err = m.InsertTx(tx, bean)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (m *MyStore) InsertTx(tx *sql.Tx, bean interface{}) error {
	fields, err := getDBFields(bean)

	if err != nil {
		tx.Rollback()
		return err
	}

	ss, err := getSqlStruct(fields)

	if err != nil {
		tx.Rollback()
		return err
	}

	sql := `insert into %s(%s) values(%s)`

	ph := make([]string, len(ss.cols))

	for i := 0; i < len(ss.cols); i++ {
		ph[i] = "?"
	}

	q := fmt.Sprintf(sql, getTableName(bean), strings.Join(ss.cols, ","), strings.Join(ph, ","))

	r, err := tx.Exec(q, ss.args...)

	if err != nil {
		tx.Rollback()
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	if id != 0 {
		err := setValue(bean, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}

func (m *MyStore) Insertx(bean ...interface{}) error {
	tx, err := m.BeginTx()

	if err != nil {
		return err
	}

	for _, item := range bean {
		if err := m.InsertTx(tx, item); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (m *MyStore) Update(bean interface{}) error {
	tx, err := m.BeginTx()

	if err != nil {
		return err
	}

	err = m.UpdateTx(tx, bean)

	if err != nil {
		return err
	}

	return m.Commit(tx)
}

func (m *MyStore) UpdateTx(tx *sql.Tx, bean interface{}) error {
	fields, err := getDBFields(bean)

	if err != nil {
		tx.Rollback()
		return err
	}

	ss, err := getSqlStruct(fields)

	if err != nil {
		tx.Rollback()
		return err
	}

	if len(ss.cols) == 0 {
		tx.Rollback()
		return errors.New("no col should be updated")
	}

	sql := "update " + getTableName(bean) + " SET "

	for i := 0; i < len(ss.cols); i++ {
		sql += ss.cols[i] + " = ?"

		if i != len(ss.cols)-1 {
			sql += ","
		}
	}

	sql += " where " + ss.pcol + " = ?"

	ss.args = append(ss.args, ss.pval)

	_, err = tx.Exec(sql, ss.args...)

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil

}

func (m *MyStore) Updatex(bean ...interface{}) error {

	tx, err := m.BeginTx()
	if err != nil {
		return err
	}

	for _, item := range bean {
		err := m.UpdateTx(tx, item)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
