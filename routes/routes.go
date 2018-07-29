package routes

import (
	"go-manager/handler/manager/admin"
	"go-manager/handler/manager/company"
	"go-manager/middleware"
	"net/http"

	mux "github.com/julienschmidt/httprouter"
)

type ManagerRouter struct {
	mux.Router
	router *mux.Router
}

func (mr *ManagerRouter) Route(method string, path string, handlerFunc http.HandlerFunc) {
	sourceHandler := http.HandlerFunc(handlerFunc)
	logHandler := middleware.LogHandler(sourceHandler)
	loginHandler := middleware.LogginHandler(logHandler)
	mr.router.Handler(method, path, loginHandler)
}

func NewRouter() *mux.Router {
	mr := ManagerRouter{router: mux.New()}
	// login
	mr.Route("POST", "/manager/login", admin.Login)
	// company
	mr.Route("POST", "/manager/company", company.Create)
	mr.Route("PUT", "/manager/company", company.Update)
	mr.Route("GET", "/manager/company/deployment", company.GetDeployment)
	// admin
	mr.Route("GET", "/manager/api/admin/signinError", admin.SigninError)

	return mr.router
}
