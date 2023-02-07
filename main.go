package main

import (
	"fmt"
	"net/http"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/cleisommais/oauth-service-v1/routes"
	"github.com/cleisommais/oauth-service-v1/db"	
)

func main() {
	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		log.Fatalf("Error loading .env file: %v", errEnvLoad)
	}	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	conn, err := db.CreatePostgresConnection()
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
	}
	defer conn.Close()	

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	r.HandleFunc("/test", routes.TestHandler).Methods("GET")

	fmt.Println("Listening on port", port)
	httpErr := http.ListenAndServe(":"+port, r)
	if httpErr != nil {
		fmt.Println(httpErr)
	}
}