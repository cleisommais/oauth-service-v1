package controllers

import (
	"net/http"

	models "github.com/cleisommais/oauth-service-v1/models"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	render := render.New()
	id := r.Header.Get("id")
	if id == "" {
		log.WithFields(log.Fields{
			"status": http.StatusBadRequest,
		}).Error("Id query parameter is required")
		render.JSON(w, http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Message: "Id query parameter is required"})
		return
	}

	login := r.Header.Get("login")
	if login == "" {
		log.WithFields(log.Fields{
			"status": http.StatusBadRequest,
		}).Error("Login query parameter is required")
		render.JSON(w, http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Message: "Login query parameter is required"})
		return
	}
	response := &models.User{Id: id, FirstName: "Cleison", LastName: "Melo", Login: login, Password: "123"}
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, http.StatusOK, response)
}
