package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/cleisommais/oauth-service-v1/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	r.HandleFunc("/test", routes.TestHandler).Methods("GET")

	fmt.Println("Listening on port", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println(err)
	}
}