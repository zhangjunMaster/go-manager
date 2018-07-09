package handler

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(1234)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Println("-------", se.Error())
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
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}
