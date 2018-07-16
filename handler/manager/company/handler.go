package company

import (
	"encoding/json"
	"fmt"
	"go-manager/handler"
	"go-manager/handler/manager/admin"
	"go-manager/handler/manager/license"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	mux "github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var company CompanyModel
	var admin admin.AdminModel
	var license license.LicenseModel
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
	if err := json.Unmarshal(body, &license); err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	company = CompanyToEntity(company)
	admin.Company_id = company.ID
	admin = AdminToEntity(admin)
	license = LicenseToEntity(license)
	license.Company_id = company.ID
	//user.Birthday = time.Now()
	fmt.Printf("%+v", company)
	_, err = CreateCompany(company, admin, license)
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	handler.HandleOk(w, company)
}

func Update(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var company CompanyModel
	var admin admin.AdminModel
	var license license.LicenseModel
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
	if err := json.Unmarshal(body, &license); err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	//company = CompanyToEntity(company)
	admin.Company_id = company.ID
	//admin = AdminToEntity(admin)
	//license = LicenseToEntity(license)
	license.Company_id = company.ID
	//user.Birthday = time.Now()

	_, err = UpdateCompany(company, admin, license)
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	handler.HandleOk(w, company)
}

func GetDeployment(w http.ResponseWriter, r *http.Request, params mux.Params) {
	domainName := params.ByName("domainName")
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

}
