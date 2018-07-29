package middleware

import (
	"encoding/json"
	"fmt"
	"go-manager/auth"
	"go-manager/handler"
	"log"
	"net/http"
	"time"
)

func LogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(" %s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// http.HandlerFunc 传参是 func(w http.ResponseWriter, r *http.Request){}
// 只要是 func(w http.ResponseWriter, r *http.Request){} 都是 type http.HandlerFunc
// type http.HandlerFunc 有个方法是 ServeHTTP,也就是 可以调用
// http.Handler 是个interface， 有个ServeHTTP 方法
// type http.HandlerFunc 继承了 http.Handler接口，所以 可以 next http.Handler，调用ServeHTTP
// 匿名函数
func LoginHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var result = make(map[string]string)
		result["name"] = "haha"
		next.ServeHTTP(w, r)
	})
}

func AuthTokenHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := auth.Read(w, r)
		value, _ := auth.ParseToken(cookie)
		fmt.Printf("%+v", value)
		next.ServeHTTP(w, r)
	})
}

func SetCookieHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/manager/login" {
			next.ServeHTTP(w, r)
			return
		}
		cookie, _ := auth.Read(w, r)
		mapCookie, err := auth.ParseToken(cookie)
		if err != nil {
			statusError := handler.StatusError{Code: 401, Err: err}
			statusError.HandleError(w)
			return
		}
		//fmt.Printf("%+v", mapCookie)
		cookieJson, _ := json.Marshal(mapCookie)
		cookieString := string(cookieJson)
		token, _ := auth.CreateToken(cookieString)
		auth.Set(w, r, token)
		fmt.Println("cookie err:", err)
		next.ServeHTTP(w, r)
	})
}
