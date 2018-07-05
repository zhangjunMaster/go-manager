package company

import (
	"encoding/json"
	"fmt"
	"go-manager/handler"
	"io"
	"io/ioutil"
	"net/http"

	mux "github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var user User
	//将信息读取到[]byte 中
	//r.Body->io.ReadCloser类型   io.LimitReader(r Reader) => Reader  ioutil.ReadAll=> []byte
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	fmt.Println("-----body:", body)
	//r.Body是否关闭
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	//user.Birthday = time.Now()
	_, err = CreateUser(user)
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
	}

	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request, _ mux.Params) {

}

func GetDeployment(w http.ResponseWriter, r *http.Request, _ mux.Params) {

}
