package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cleisommais/oauth-service-v1/db"
	"github.com/cleisommais/oauth-service-v1/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
)

const (
	LocalEnv   = "LOCAL"
	DefaultPort = "8000"
)

func init() {
	if LocalEnv == strings.ToUpper(os.Getenv("ENV")) {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	dbConn, err := db.CreatePostgresConnection()
	if err != nil {
		logrus.WithError(err).Fatal("Error connecting to Postgres")
	}
	defer dbConn.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})
	r.HandleFunc("/test", routes.TestHandler).Methods("GET")

	logrus.WithField("port", port).Info("Listening on port")
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		logrus.WithError(err).Fatal("Error starting the server")
	}
}
