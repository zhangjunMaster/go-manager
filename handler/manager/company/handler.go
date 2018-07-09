package company

import (
	"encoding/json"
	"fmt"
	"go-manager/handler"
	"go-manager/handler/manager/admin"
	"io"
	"io/ioutil"
	"net/http"

	mux "github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var company CompanyModel
	var admin admin.AdminModel
	//将信息读取到[]byte 中
	//r.Body->io.ReadCloser类型   io.LimitReader(r Reader) => Reader  ioutil.ReadAll=> []byte
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
	if err := json.Unmarshal(body, &company); err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	if err := json.Unmarshal(body, &admin); err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	company = CompanyToEntity(company)
	admin = AdminToEntity(admin)
	//user.Birthday = time.Now()
	fmt.Printf("%+v", company)
	_, err = CreateCompany(company)
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	handler.HandleOk(w, company)
}

func Update(w http.ResponseWriter, r *http.Request, _ mux.Params) {

}

func GetDeployment(w http.ResponseWriter, r *http.Request, _ mux.Params) {

}
