package middleware

import (
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
func LogginHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start := time.Now()
		//log.Printf("Started %s %s", r.Method, r.URL.Path)
		var result = make(map[string]string)
		result["name"] = "haha"
		next.ServeHTTP(w, r)
		//log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
