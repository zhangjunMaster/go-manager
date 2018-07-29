package admin

import (
	"go-manager/handler"
	"go-manager/lib"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	//fmt.Printf("%+v", r.Method)
	//fmt.Println(r.URL)
	username, _ := lib.DecodeBase64(header.Get("username"))
	password, _ := lib.DecodeBase64(header.Get("password"))
	encryptoPassword := lib.Md5Salt(password)
	//fmt.Println(encryptoPassword)
	rows, err := Get(username, encryptoPassword)
	if err != nil {
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	admin := AdminToModel(rows[0])
	handler.HandleOk(w, admin)
}

func SigninError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "forbitten", 401)
}
