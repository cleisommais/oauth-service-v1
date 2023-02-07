// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth", AuthController.Authenticate).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
