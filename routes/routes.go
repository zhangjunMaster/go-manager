package routes

import (
	"go-manager/handler/manager/admin"
	"go-manager/handler/manager/company"
	"go-manager/handler/manager/dashbord"
	"go-manager/handler/manager/department"
	"go-manager/handler/manager/user"
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
	// 日志
	logHandler := middleware.LogHandler(sourceHandler)
	// 设置cookie
	//SetCookieHandler := middleware.SetCookieHandler(logHandler)
	// 登录
	//loginHandler := middleware.LoginHandler(SetCookieHandler)
	mr.router.Handler(method, path, logHandler)
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
	// dashbord(未填充)
	mr.Route("GET", "/manager/dashboard/:id/user/activated", dashbord.GetActivatedUserCount)
	mr.Route("GET", "/manager/dashboard/:id/device", dashbord.GetActivatedDeviceCount)
	mr.Route("GET", "/manager/dashboard/:id/user/active", dashbord.GetActiveUserData)
	mr.Route("GET", "/manager/dashboard/:id/os", dashbord.GetOSOfPCData)
	mr.Route("GET", "/manager/dashboard/:id/user/online", dashbord.GetOnlineUserData)
	// user
	mr.Route("POST", "/manager/user", user.Create)
	mr.Route("DELETE", "/manager/user", user.Delete)
	mr.Route("PUT", "/manager/user", user.Update)
	mr.Route("GET", "/manager/user/all", user.GetAllUsers)
	mr.Route("GET", "/manager/user/baseinfo/:id", user.GetOneUser)
	mr.Route("GET", "/manager/user/strategies/all/:id", user.GetAllStrategies)
	mr.Route("GET", "/manager/user/strategies/active", user.GetActiveStrategy)
	mr.Route("GET", "/manager/user/application", user.GetApplications)
	mr.Route("GET", "/manager/user/devices/:id", user.GetDevices)
	mr.Route("PUT", "/manager/user/password", user.ChangePwd)
	mr.Route("POST", "/manager/user/active", user.Active)
	// department
	mr.Route("POST", "/manager/department", department.Create)
	mr.Route("DELETE", "/manager/department", department.Delete)
	mr.Route("PUT", "/manager/department", department.Update)
	mr.Route("GET", "/manager/department/all", department.GetAllDepartments)
	mr.Route("GET", "/manager/department/one", department.GetOneDepartment)
	mr.Route("GET", "/manager/department/portion", department.GetChildrenDepartments)

	return mr.router
}
