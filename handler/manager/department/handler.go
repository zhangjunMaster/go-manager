package department

import (
	"encoding/json"
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
func GetOneDepartment(w http.ResponseWriter, r *http.Request) {}
func GetChildrenDepartments(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	departmentId := vars["departmentId"][0]
	condition := vars["condition"][0]
	count := vars["count"][0]
	start := vars["start"][0]
	companyId := vars["companyId"][0]
	rows, err := GetChildrenModel(companyId, departmentId, condition, count, start)
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	handler.HandleOk(w, rows)
}
