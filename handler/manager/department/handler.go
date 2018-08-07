package department

import (
	"encoding/json"
	"fmt"
	"go-manager/handler"
	"io"
	"io/ioutil"
	"net/http"
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
	var companyId string
	//companyId := vars["companyId"]
	if len(vars) == 0 {
		companyId = "a3c69b85-6745-436e-a1c1-f99d42e6f0eb"
	}
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
func GetOneDepartment(w http.ResponseWriter, r *http.Request) {}
func GetChildrenDepartments(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	var (
		departmentId string
		condition    string
		count        string
		start        string
	)
	if len(vars) > 0 {
		departmentId = vars["departmentId"][0]
		condition = vars["condition"][0]
		count = vars["count"][0]
		start = vars["start"][0]
	}
	fmt.Println(vars)
}
