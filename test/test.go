package main

import (
	"fmt"
	"go-manager/model"
)

func main() {
	var companyModel = model.Model{TableName: "company"}
	var mdb = companyModel.Open()

	rows2, _ := mdb.DB.Query("select * from company")
	//返回所有列的名字[]string,列名是字符串
	cols, _ := rows2.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据，scans是一个slice
	//长度是列的个数
	scans := make([]interface{}, len(cols))
	//这里scans是列名的地址
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	fmt.Println("--scans:", scans)
	i := 0
	var result []map[string]string
	for rows2.Next() {
		//填充数据 Query的结果是Rows，方法func (rs *Rows) Scan(dest ...interface{}) error
		//将结果填入到scans中,scans中的是指针，取得是vals中的指针
		rows2.Scan(scans...)
		fmt.Println("----scans:", &scans[0])
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result = append(result, row)
		i++
	}
	fmt.Printf("%+v", result)
}
