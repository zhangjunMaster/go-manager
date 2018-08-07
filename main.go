package main

import (
	"go-manager/routes"
	"log"
	"net/http"
	//_ "net/http/pprof"
	"github.com/feixiao/httpprof"
)

func main() {
	router := routes.NewRouter()
	router = httpprof.WrapRouter(router)
	log.Fatal(http.ListenAndServe(":8091", router))
}
