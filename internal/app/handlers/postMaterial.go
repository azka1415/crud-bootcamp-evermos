package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/exceptions"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/responses"
	log "github.com/sirupsen/logrus"
)

func HandlePostMaterial(w http.ResponseWriter, r *http.Request) {
	handleLogger := log.WithFields(log.Fields{"post": "/materials"})

	var updatedMaterial models.UpdateMaterial
	err := json.NewDecoder(r.Body).Decode(&updatedMaterial)

	if err != nil {
		exceptions.BadBodyException(w)
		handleLogger.Error(err)
		return
	}
	matService := models.NewMaterialService()

	mat, err := matService.NewMaterial(updatedMaterial)

	if err != nil {
		exceptions.BadBodyException(w)
		handleLogger.Error(err)
		return
	}

	responses.NewMaterialResponse(w, mat)
}
