package model

import (
	"database/sql"
	"fmt"
	"go-manager/lib"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root"           //用户名
	dbpassword = "123456"         //密码
	dbname     = "rdc_manager"    //表名
)

type Model struct {
	db        *sql.DB
	TableName string
}

func (M *Model) Open() *Model {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_manager")
	//如下是创建pool
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
	//如上创建pool
	if err != nil {
		panic(err)
	}
	M.db = db
	return M
}

func (M *Model) Create(v interface{}) (sql.Result, error) {
	var placeholders []string
	colNames := lib.QuetoKey(v)
	colValues := lib.QuetoValue(v)
	for _, _ = range colNames {
		placeholders = append(placeholders, "?")
	}
	sqlStr := fmt.Sprintf("INSERT INTO %s (%v) VALUES (%v)",
		M.TableName,
		strings.Join(colNames, ", "),
		strings.Join(placeholders, ", "))
	fmt.Println(sqlStr, colValues)
	res, err := M.db.Exec(sqlStr, colValues...)
	defer M.db.Close()
	return res, err
}
