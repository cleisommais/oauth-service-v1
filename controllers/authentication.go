package controllers

import (
	"net/http"
	"time"

	"github.com/cleisommais/oauth-service-v1/models"
	"github.com/palantir/stacktrace"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	ren := render.New()
	log.WithFields(log.Fields{
		"status": http.StatusOK,
	}).Info("Login requested")
	response := &models.User{Id: "21313", FirstName: "Cleison", LastName: "Melo", Login: "cleison", Password: "123", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	w.Header().Set("Content-Type", "application/json")
	if err := ren.JSON(w, http.StatusOK, response); err != nil {
		stacktrace.Propagate(err, "Failed to encode response to JSON")
		log.WithError(err).Error("Failed to encode response to JSON")
		ren.JSON(w, http.StatusInternalServerError, map[string]string{"msg": "Internal server error"})
	}
}
