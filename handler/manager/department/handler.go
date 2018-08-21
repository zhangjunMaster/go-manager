package department

import (
	"encoding/json"
	"fmt"
	"go-manager/handler"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var department DepartmentModel
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	//r.Body是否关闭
	if err := r.Body.Close(); err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	if err := json.Unmarshal(body, &department); err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	department = DepartmentToEntity(department)

	ids := []string{department.Parent_department_id}
	rows, _ := GetOneModel(ids)
	department.Full_department_id = rows[0]["full_department_id"] + "/" + department.ID
	fmt.Printf("%v", department)
	_, err = CreateDepartment(department)

	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	handler.HandleOk(w, department)
}
func Delete(w http.ResponseWriter, r *http.Request) {}
func Update(w http.ResponseWriter, r *http.Request) {}
func GetAllDepartments(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	companyId := vars["companyId"][0]
	rows, err := GetAllModel(companyId)
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	results := make([]map[string]interface{}, len(rows))
	for _, value := range rows {
		if value != nil {
			results = append(results, DepartToModel(value))
		}
	}
	results = CreateTree(results, "0")
	handler.HandleOk(w, results)
}

// get one department
func GetOneDepartment(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars["id"][0]
	ids := []string{id}
	rows, err := GetOneModel(ids)
	if err != nil {
		fmt.Println(err)
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	data := DepartToModel(rows[0])

	if data["parentDepartmentId"] == "0" {
		data["fullDepartmentName"] = data["name"]
		handler.HandleOk(w, data)
		return
	} else {
		fullDepartmentId := data["fullDepartmentId"].(string)
		ids := strings.Split(fullDepartmentId, "/")
		rows, err := GetOneModel(ids)
		if err != nil {
			fmt.Println(err)
			statusError := handler.StatusError{Code: 500, Err: err}
			statusError.HandleError(w)
			return
		}
		data["fullDepartmentName"] = CreatFullDepartmentName(rows)
	}

	handler.HandleOk(w, data)
}

// get children department
func GetChildrenDepartments(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	departmentId := vars["departmentId"][0]
	condition := vars["condition"][0]
	count := vars["count"][0]
	start := vars["start"][0]
	companyId := vars["companyId"][0]
	result, err := GetChildrenModel(companyId, departmentId, condition, count, start)
	rows := result["data"].(map[int]map[string]string)
	fmt.Println("------rows:", len(rows))
	results := make([]map[string]interface{}, len(rows))
	for index, value := range rows {
		fmt.Printf("%v", value)
		if value != nil {
			results[index] = DepartToModel(value)
		}
	}
	result["departments"] = results
	delete(result, "data")
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	handler.HandleOk(w, result)
}
