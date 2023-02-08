package controllers

import (
	"net/http"
	"strconv"

	models "github.com/cleisommais/oauth-service-v1/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	render := render.New()
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		log.WithFields(log.Fields{
			"status": http.StatusBadRequest,
		}).Error("Id query parameter is required")
		render.JSON(w, http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Message: "Id query parameter is required"})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.WithFields(log.Fields{
			"status": http.StatusBadRequest,
		}).Error("Invalid id query parameter")
		render.JSON(w, http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Message: "Invalid id query parameter"})
		return
	}
	login, ok := vars["login"]
	if !ok {
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
