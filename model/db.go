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

type Transaction struct {
	Sql    string
	Values []interface{}
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
	colNames := lib.QuoteKey(v)
	colValues := lib.QuoteValue(v)
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

func (M *Model) QuoteUpdateFields(v interface{}) (string, []interface{}) {
	var placeholders []string
	var updateValues []interface{}
	colNames := lib.QuoteKey(v)
	colValues := lib.QuoteValue(v)
	for index, key := range colNames {
		if colValues[index] != nil && key != "ID" && colValues[index] != "" && colValues[index] != 0 && colValues[index] != "0001-01-01 00:00:00" {
			sqlStr := fmt.Sprintf("%v=%v",
				key,
				"?")
			updateValues = append(updateValues, colValues[index])
			placeholders = append(placeholders, sqlStr)
		}
	}
	fmt.Println("----placeholders", placeholders)
	fmt.Println("----updateValues", updateValues)

	colNameString := strings.Join(placeholders, ", ")
	return colNameString, updateValues
}

// 回滚
func clearTransaction(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		return err
	}
	return nil
}

func (M *Model) Transaction(transaction []Transaction) error {
	tx, err := M.db.Begin()
	if err != nil {
		return err
	}
	defer M.db.Close()
	defer clearTransaction(tx)
	for _, t := range transaction {
		stmt, err := tx.Prepare(t.Sql)
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = stmt.Exec(t.Values...)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
