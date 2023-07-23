package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/exceptions"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/responses"
	"github.com/azka1415/crud-bootcamp-evermos/tools/utils"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func HandleUpdateMaterial(w http.ResponseWriter, r *http.Request) {
	materialID, err := utils.ConvertToInt(chi.URLParam(r, "id"))
	if err != nil {
		exceptions.BadParamException(w)
		log.Error("invalid ID")
		return
	}

	handleLogger := log.WithFields(log.Fields{"put": fmt.Sprintf("/materials/%v", materialID)})

	matService := models.NewMaterialService()

	exists, err := matService.MaterialExistsByID(materialID)

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

	var updatedMaterial models.UpdateMaterial
	err = json.NewDecoder(r.Body).Decode(&updatedMaterial)
	if err != nil {
		exceptions.BadBodyException(w)
		handleLogger.Error(err)
		return
	}

	mat, err := matService.UpdateMaterial(materialID, updatedMaterial)

	if err != nil {
		exceptions.BadRequestException(w, err)
		handleLogger.Error(err)
		return
	}

	responses.UpdateMaterialResponse(w, mat)

}
