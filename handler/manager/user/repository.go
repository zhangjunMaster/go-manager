package user

import (
	"database/sql"
	"fmt"
	"go-manager/lib"
	"go-manager/model"
	"strconv"
)

const (
	tableName   = "user"
	timeFormart = "2006-01-02 15:04:05"
)

// 将客户端传来的数据转换成sql的字段
type UserModel struct {
	ID            string       `json:"id"`
	Company_id    string       `json:"companyId"`
	Name          string       `json:"name"`
	Department_id string       `json:"departmentId"`
	Source        int          `json:"source"`
	Password      string       `json:"password"`
	Email         string       `json:"email"`
	Mobile        string       `json:"mobile"`
	Login_name    string       `json:"loginName"`
	Title         string       `json:"title"`
	Avatar_path   string       `json:"avatarPath"`
	Create_date   lib.JsonTime `json:"createDate"`
	Last_update   lib.JsonTime `json:"lastUpdate"`
}

var userModel = model.Model{TableName: tableName}

func CreateUserModel(user *UserModel) (sql.Result, error) {
	var mdb = userModel.Open()
	//fmt.Printf("%+v", &user)
	result, err := mdb.Create(*user)
	return result, err
}

func GetAllModel(params map[string]string) (map[string]interface{}, error) {
	var mdb = userModel.Open()
	var result = make(map[string]interface{})
	departmentId := params["departmentId"]
	companyId := params["companyId"]
	count, _ := strconv.Atoi(params["count"])
	condition := params["condition"]
	start, _ := strconv.Atoi(params["start"])
	start = (start - 1) * count
	queryValues := []interface{}{companyId, departmentId}
	queryDepartSql := `
                         SELECT t1.*, t2.full_department_id
                         FROM user t1 , department t2
                         WHERE t1.company_id = ?
                         AND t1.department_id = ?
                         AND t1.is_deleted = 0
                         `
	if condition != "" {
		queryDepartSql += `AND t1.name LIKE ?`
		queryValues = append(queryValues, condition)
	}
	queryDepartSql += `
                      AND t2.id = ?
                      ORDER BY create_date DESC limit ?,?
					  `
	queryValues = append(queryValues, departmentId, start, count)
	countQueryValues := []interface{}{companyId, departmentId}
	countQuerySql := `
                        SELECT COUNT(*) AS total
                        FROM user t1
                        WHERE t1.company_id = ?
                        AND t1.department_id = ?
                        AND t1.is_deleted = 0
                        `
	if condition != "" {
		countQuerySql += `AND t1.name LIKE ?`
		countQueryValues = append(countQueryValues, condition)
	}

	fmt.Println(queryDepartSql, queryValues)

	rows, err := mdb.Query(queryDepartSql, queryValues)
	fmt.Println(len(rows))
	countRows, err := mdb.Query(countQuerySql, countQueryValues)
	result["total"] = countRows[0]["total"]
	result["start"] = start
	result["data"] = rows
	return result, err
}
