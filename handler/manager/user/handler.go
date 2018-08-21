package user

import (
	"encoding/json"
	"go-manager/handler"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user UserModel
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	if err := r.Body.Close(); err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	userData := UserToEntity(&user)
	result, err := CreateUserModel(userData)
	log.Println(result)
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	handler.HandleOk(w, user)
}
func Update(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {}
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	params := make(map[string]string)
	params["companyId"] = query["companyId"][0]
	params["departmentId"] = query["departmentId"][0]
	params["condition"] = query["condition"][0]
	params["count"] = query["count"][0]
	params["start"] = query["start"][0]
	result, err := GetAllModel(params)
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	rows := result["data"].(map[int]map[string]string)
	results := make([]map[string]interface{}, len(rows))
	for index, value := range rows {
		if value != nil {
			results[index] = UserToModel(value)
		}
	}
	result["users"] = results
	delete(result, "data")
	handler.HandleOk(w, result)
}
func GetOneUser(w http.ResponseWriter, r *http.Request)        {}
func GetAllStrategies(w http.ResponseWriter, r *http.Request)  {}
func GetActiveStrategy(w http.ResponseWriter, r *http.Request) {}
func GetApplications(w http.ResponseWriter, r *http.Request)   {}
func GetDevices(w http.ResponseWriter, r *http.Request)        {}
func ChangePwd(w http.ResponseWriter, r *http.Request)         {}
func Active(w http.ResponseWriter, r *http.Request)            {}
