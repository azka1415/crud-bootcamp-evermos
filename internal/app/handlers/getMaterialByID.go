package handlers

import (
	"fmt"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/exceptions"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/responses"
	"github.com/azka1415/crud-bootcamp-evermos/tools/utils"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func HandleGetMaterialByID(w http.ResponseWriter, r *http.Request) {
	materialID, err := utils.ConvertToInt(chi.URLParam(r, "id"))
	if err != nil {
		exceptions.BadParamException(w)
		log.Error("invalid ID")
		return
	}

	handleLogger := log.WithFields(log.Fields{"get": fmt.Sprintf("/materials/%v", materialID)})
	materialService := models.NewMaterialService()

	exists, err := materialService.MaterialExistsByID(materialID)

	if err != nil {
		exceptions.NotFoundException(w, err)
		handleLogger.Error(err)
		return
	}

	if !exists {
		exceptions.NotFoundException(w, err)
		handleLogger.Error(err)
		return
	}

	m, err := materialService.GetByID(materialID)

	if err != nil {
		exceptions.NotFoundException(w, err)
		handleLogger.Error(err)
		return
	}

	responses.GetMaterialByIDResponse(w, m)
	handleLogger.Info("Get Material By ID success")
}
