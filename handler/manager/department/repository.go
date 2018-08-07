package department

import (
	"database/sql"
	"go-manager/lib"
	"go-manager/model"
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

func GetChildrenModel(departmentId string, condition string, count string, start string) {}
func CalculateUCountByDId(departmentIds []map[int]string)                                {}
