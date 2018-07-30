package handler

import (
	"encoding/json"
	"net/http"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code    int    `json:"code"`
	Err     error  `json:"data"`
	Message string `json:"message"`
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

func (se StatusError) ErrJson() {

}

//错误处理
func (se StatusError) HandleError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//http.Error(w, se.Error(), se.Status())
	if err := json.NewEncoder(w).Encode(se); err != nil {
		panic(err)
	}
	return
	//panic(se.Error())
}

//正确返回
func HandleOk(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//http.Error(w, se.Error(), se.Status())
	data := make(map[string]interface{})
	data["errCode"] = "0"
	data["data"] = v
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
