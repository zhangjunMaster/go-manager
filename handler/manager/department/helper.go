package department

import (
	"fmt"
	"go-manager/lib"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func DepartmentToEntity(department DepartmentModel) DepartmentModel {
	if department.ID == "" {
		uid := uuid.Must(uuid.NewV4())
		//fmt.Println("uid:", fmt.Sprintf("%s", uid))
		department.ID = fmt.Sprintf("%s", uid)
	}

	department.Create_date = lib.JsonTime(time.Now())
	department.Last_update = lib.JsonTime(time.Now())
	return department
}

func DepartToModel(source map[string]string) map[string]interface{} {
	target := make(map[string]interface{})
	target["companyId"] = source["company_id"]
	target["fullDepartmentId"] = source["full_department_id"]
	target["id"] = source["id"]
	target["incrementalOperation"] = source["incremental_operation"]
	target["isImGroupCreated"] = source["is_im_group_created"]
	target["lastUpdate"] = source["last_update"]
	target["name"] = source["name"]
	target["parentDepartmentId"] = source["parent_department_id"]
	target["parentDepartmentName"] = source["parent_department_name"]
	target["source"] = false
	return target
}

func fn(data []map[string]interface{}, parentDepartmentId string) []map[string]interface{} {
	var result []map[string]interface{}
	var temp []map[string]interface{}
	for _, value := range data {
		if value["parentDepartmentId"] == parentDepartmentId {
			result = append(result, value)
			id := value["id"]
			temp = fn(data, id.(string))
			if len(temp) > 0 {
				value["children"] = temp
			} else {
				value["children"] = make([]map[string]string, 0)
			}
		}
	}
	return result
}

func CreateTree(departments []map[string]interface{}, pid string) []map[string]interface{} {
	if pid == "" {
		pid = "0"
	}
	result := fn(departments, pid)
	return result
}

func CreatFullDepartmentName(departments map[int]map[string]string) string {
	fullDepartmentName := ""
	for _, v := range departments {
		fullDepartmentName += (v["name"] + "/")
	}
	fullDepartmentName = strings.Replace(fullDepartmentName, `\/$`, "", 1)
	return fullDepartmentName
}
