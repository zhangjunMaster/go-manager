package dashbord

import (
	"fmt"
	"go-manager/handler"
	"net/http"
	"strings"
)

type Param struct {
	Key   string
	Value string
}

type key string

var paramKey key = "id"
var p *Param

func GetActivatedUserCount(w http.ResponseWriter, r *http.Request) {
	s := strings.TrimPrefix(r.URL.Path, "/manager/dashboard/")
	id := strings.TrimRight(s, "user/activated")
	rows, err := GetActivatedUserCountData(id)
	if err != nil {
		fmt.Printf("%+v", err)
		statusError := handler.StatusError{Code: 500, Err: err}
		statusError.HandleError(w)
		return
	}
	handler.HandleOk(w, rows[0])
}

func GetActivatedDeviceCount(w http.ResponseWriter, r *http.Request) {

}
func GetActiveUserData(w http.ResponseWriter, r *http.Request) {

}
func GetOSOfPCData(w http.ResponseWriter, r *http.Request) {

}
func GetOnlineUserData(w http.ResponseWriter, r *http.Request) {

}
