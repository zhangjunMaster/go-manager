package department

import (
	"database/sql"
	"fmt"
	"go-manager/lib"
	"go-manager/model"
	"strconv"
)

const (
	tableName   = "department"
	timeFormart = "2006-01-02 15:04:05"
)

// 将客户端传来的数据转换成sql的字段
type DepartmentModel struct {
	ID                       string       `json:"id"`
	Company_id               string       `json:"companyId"`
	Department_id            string       `json:"name"`
	Name                     string       `json:"domainName"`
	Description              string       `json:"managerServer"`
	Parent_department_id     string       `json:"isPrivateDevelopment"`
	Source                   int          `json:"edition"`
	Is_im_group_created      int          `json:"type"`
	Full_department_id       string       `json:"originalManagerServer"`
	Incremental_operation    int          `json:"isDeleted"`
	User_device_num          int          `json:"userDeviceNum"`
	Is_limit_user_device_num int          `json:"isLimitUserDeviceNum"`
	Create_date              lib.JsonTime `json:"createDate"`
	Last_update              lib.JsonTime `json:"lastUpdate"`
	Create_date_at_hub       lib.JsonTime `json:"createDateAtHub"`
}

var departmentModel = model.Model{TableName: tableName}

func CreateDepartment(departmentData DepartmentModel) (sql.Result, error) {
	var mdb = departmentModel.Open()
	result, err := mdb.Create(departmentData)
	return result, err
}

/**
* The operation is get all departments
 */
func GetAllModel(companyId string) (map[int]map[string]string, error) {
	var mdb = departmentModel.Open()
	params := []interface{}{companyId}
	queryDepartSql := `
                         SELECT t1.*,t2.name AS parent_department_name
                         FROM department t1
                         inner join department t2
                         on t1.parent_department_id = t2.id
						 WHERE t1.company_id = ?
						 AND t1.parent_department_id <> "0"
                         `
	queryTopSql := `
                      SELECT t1.*,t1.name AS parent_department_name
                      FROM department t1
                      WHERE t1.company_id = ?
                      AND t1.parent_department_id = "0"
                      `
	rows, err := mdb.Query(queryDepartSql, params)
	length := len(rows)
	topRows, err := mdb.Query(queryTopSql, params)
	rows[length] = topRows[0]
	return rows, err
}

func GetChildrenModel(companyId string, departmentId string, condition string, c string, start string) (map[string]interface{}, error) {
	var mdb = departmentModel.Open()
	var result = make(map[string]interface{})
	count, _ := strconv.Atoi(c)
	begin, _ := strconv.Atoi(start)
	begin = (begin - 1) * count
	departmentName := condition
	params := []interface{}{companyId, departmentId}
	countParams := []interface{}{companyId, departmentId}
	queryChildrenDepartSql := `
                         SELECT t1.*
                         FROM department t1
                         WHERE t1.company_id = ?
                         AND t1.parent_department_id = ?
                         `
	if departmentName != "" {
		queryChildrenDepartSql += `AND t1.name LIKE "%${departmentName}%"`
	}
	queryChildrenDepartSql += ` 
							  ORDER BY t1.create_date DESC limit ?,?
							  `
	params = append(params, begin, count)

	countQuerySql := `
                     SELECT COUNT(*) AS total FROM department t1
                     WHERE t1.company_id = ?
                     AND t1.parent_department_id = ?
					 `
	if departmentName != "" {
		countQuerySql += `AND t1.name LIKE "%${departmentName}%"`
	}
	rows, err := mdb.Query(queryChildrenDepartSql, params)
	//length := len(rows)
	fmt.Println("row length:", len(rows))
	fmt.Printf("%v", queryChildrenDepartSql)
	fmt.Printf("%v", params)

	countRows, err := mdb.Query(countQuerySql, countParams)
	result["total"] = countRows[0]["total"]
	result["start"] = start
	result["data"] = rows
	return result, err
}
func CalculateUCountByDId(departmentIds []map[int]string) {}

func GetOneModel(ids []string) (map[int]map[string]string, error) {
	var mdb = departmentModel.Open()
	params := make([]interface{}, len(ids))
	//params = copy(params, ids[:])
	for index, v := range ids {
		params[index] = v
	}
	sql := `
			SELECT * FROM department
			WHERE id IN (?)
			ORDER BY create_date
			`
	fmt.Println(sql)
	fmt.Println(params)
	rows, err := mdb.Query(sql, params)
	return rows, err
}
