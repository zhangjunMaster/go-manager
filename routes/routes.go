package routes

import (
	"go-manager/handler/manager/company"

	mux "github.com/julienschmidt/httprouter"
)

func NewRouter() *mux.Router {

	router := mux.New()
	//company
	router.POST("/manager/company", company.Create)
	router.PUT("/manager/company", company.Update)
	router.GET("/manager/company/deployment", company.GetDeployment)
	return router
}
