package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/exceptions"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
	"github.com/azka1415/crud-bootcamp-evermos/tools/utils"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func HandleGetMaterialByID(w http.ResponseWriter, r *http.Request) {
	materialID, err := utils.ConvertToInt(chi.URLParam(r, "id"))
	if err != nil {
		exceptions.BadRequestException(w, err)
		log.Error("invalid ID")
		return
	}

	handleLogger := log.WithFields(log.Fields{"get": fmt.Sprintf("/materials/%v", materialID)})

	var updatedMaterial models.Material
	err = json.NewDecoder(r.Body).Decode(&updatedMaterial)
	if err != nil {
		exceptions.BadRequestException(w, err)
		handleLogger.Error(err)
		return
	}
}
