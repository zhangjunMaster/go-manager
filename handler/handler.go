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

func (se StatusError) HandleError(w http.ResponseWriter) {
	fmt.Println(1234)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Println("-------", se.Error())
	//http.Error(w, se.Error(), se.Status())
	if err := json.NewEncoder(w).Encode(se); err != nil {
		panic(err)
	}
}
