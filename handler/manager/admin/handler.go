package admin

import (
	"encoding/json"
	"go-manager/auth"
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
	// 设置token,将cookie转为json,然后转为string
	adminJson, _ := json.Marshal(admin)
	adminString := string(adminJson)

	token, _ := auth.CreateToken(adminString)
	auth.Set(w, r, token)

	handler.HandleOk(w, admin)
}

func SigninError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "forbitten", 401)
}
