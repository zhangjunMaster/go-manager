package main

import (
	"log"
	"net/http"

	"go-manager/routes"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8091", router))
}
