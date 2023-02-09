package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cleisommais/oauth-service-v1/db"
	"github.com/cleisommais/oauth-service-v1/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

const (
	LOCAL = "local"
)

func start() {
	var formatter log.Formatter
	logLevel := log.InfoLevel

	if os.Getenv("ENV") == LOCAL {
		formatter = &log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
			ForceColors: true,
		}
		logLevel = log.DebugLevel
	} else {
		formatter = &log.JSONFormatter{}
	}

	log.SetFormatter(formatter)
	log.SetLevel(logLevel)
	log.Info("Running as " + os.Getenv("ENV") + " Environment")
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.WithError(err).Fatal("Error loading .env file")
	}
	start()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	dbConn, err := db.CreatePostgresConnection()
	if err != nil {
		log.WithError(err).Fatal("Error connecting to Postgres")
	}
	defer dbConn.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})
	for _, route := range routes.RoutesSetup {
		handler := routes.MakeHandler(http.HandlerFunc(route.HandlerFunc))
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)
	n.Run(":" + port)
}
