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
	DB        *sql.DB
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
	M.DB = db
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
	res, err := M.DB.Exec(sqlStr, colValues...)
	defer M.DB.Close()
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
	tx, err := M.DB.Begin()
	if err != nil {
		return err
	}
	defer M.DB.Close()
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

/**
  rows, err := db.Query("SELECT ...")
  ...
  defer rows.Close()
  for rows.Next() {
	  // 定义的接收值的变量
      var id int
	  var name string
	  // 用地址接收值
      err = rows.Scan(&id, &name)
      ...
  }
  err = rows.Err() // get any error encountered during iteration
  ...
*/

func (M *Model) Query(sql string, params []interface{}) (map[int]map[string]string, error) {
	//1.查询数据，取字段
	rows2, err := M.DB.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	//2.返回所有列的名字 []string
	cols, err := rows2.Columns()
	if err != nil {
		return nil, err
	}
	//3.一行中所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//4.获取这一行中所有列的值的地址，并写入到scans中
	scans := make([]interface{}, len(cols))
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		//填充数据 Query的结果是Rows，方法func (rs *Rows) Scan(dest ...interface{}) error
		//5.将结果填入到scans中的地址上
		rows2.Scan(scans...)
		//6.定义每行数据的格式
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	fmt.Printf("%+v", result)
	return result, nil
}
