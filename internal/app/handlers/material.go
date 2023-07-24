package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/exceptions"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/responses"
	"github.com/azka1415/crud-bootcamp-evermos/tools/utils"
	"github.com/azka1415/crud-bootcamp-evermos/tools/utils/enums/sort"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type MaterialHandler struct{}

func NewMaterialHandler() *MaterialHandler {
	return &MaterialHandler{}
}

func (h *MaterialHandler) DeleteMaterial(w http.ResponseWriter, r *http.Request) {
	materialID, err := utils.ConvertToInt(chi.URLParam(r, "id"))
	if err != nil {
		exceptions.BadParamException(w)
		log.Error("invalid ID")
		return
	}
	handleLogger := log.WithFields(log.Fields{"delete": fmt.Sprintf("/materials/%v", materialID)})
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
	_, err = matService.DeleteMaterial(materialID)

	if err != nil {
		exceptions.BadRequestException(w, err)
		handleLogger.Error(err)
		return
	}

	responses.NoContentResponse(w)
}

func (h *MaterialHandler) GetMaterial(w http.ResponseWriter, r *http.Request) {

	handleLogger := log.WithFields(log.Fields{"get": "/materials"})

	page, err := utils.ConvertToInt(utils.ParseQueryParams(r, "page"))
	if err != nil {
		err = errors.New("invalid page query param")
		handleLogger.Error(err)
		exceptions.BadQueryException(w, err)
		return
	}

	limit, err := utils.ConvertToInt(utils.ParseQueryParams(r, "limit"))
	if err != nil {
		err = errors.New("invalid limit query param")
		handleLogger.Error(err)
		exceptions.BadQueryException(w, err)
		return
	}

	sort := sort.GetSortDirection(utils.ParseQueryParams(r, "sort"))
	field := utils.ParseQueryParams(r, "field")

	materialService := models.NewMaterialService()

	m, err := materialService.GetAll(limit, page, sort, field)

	if err != nil {
		handleLogger.Error(err)
		exceptions.BadQueryException(w, err)
		return
	}

	responses.GetAllResponse(w, m, limit, page)
	handleLogger.Info("Get All Materials Success")
}

func (h *MaterialHandler) GetMaterialByID(w http.ResponseWriter, r *http.Request) {
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

	responses.GetByIDResponse(w, m)
	handleLogger.Info("Get Material By ID success")
}

func (h *MaterialHandler) PostMaterial(w http.ResponseWriter, r *http.Request) {
	handleLogger := log.WithFields(log.Fields{"post": "/materials"})

	var updatedMaterial models.PayloadMaterial
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

	responses.CreatedResponse(w, mat)
}

func (h *MaterialHandler) UpdateMaterial(w http.ResponseWriter, r *http.Request) {
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

	var updatedMaterial models.PayloadMaterial
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

	responses.UpdateByIDResponse(w, mat)
	handleLogger.Info("Updated Material success")
}
