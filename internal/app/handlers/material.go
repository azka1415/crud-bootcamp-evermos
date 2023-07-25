package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/exceptions"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/repository"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/responses"
	"github.com/azka1415/crud-bootcamp-evermos/tools/utils"
	"github.com/azka1415/crud-bootcamp-evermos/tools/utils/enums/sort"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type MaterialHandlerImpl struct {
	service models.MaterialService
}

type MaterialHandler interface {
	DeleteMaterial(w http.ResponseWriter, r *http.Request)
	GetMaterial(w http.ResponseWriter, r *http.Request)
	GetMaterialByID(w http.ResponseWriter, r *http.Request)
	PostMaterial(w http.ResponseWriter, r *http.Request)
	UpdateMaterial(w http.ResponseWriter, r *http.Request)
}

func NewMaterialHandler(service models.MaterialService) MaterialHandler {
	mh := MaterialHandlerImpl{service: service}
	var mathandle MaterialHandler = &mh

	return mathandle
}

func (h *MaterialHandlerImpl) DeleteMaterial(w http.ResponseWriter, r *http.Request) {
	materialID, err := utils.ConvertToInt(chi.URLParam(r, "id"))
	if err != nil {
		exceptions.BadParamException(w)
		log.Error("invalid ID")
		return
	}
	handleLogger := log.WithFields(log.Fields{"delete": fmt.Sprintf("/materials/%v", materialID)})

	exists, err := h.service.MaterialExistsByID(materialID)

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
	_, err = h.service.DeleteMaterial(materialID)

	if err != nil {
		exceptions.BadRequestException(w, err)
		handleLogger.Error(err)
		return
	}

	responses.NoContentResponse(w)
}

func (h *MaterialHandlerImpl) GetMaterial(w http.ResponseWriter, r *http.Request) {

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
	field := utils.CheckFieldQuery(utils.ParseQueryParams(r, "field"))
	filterTeacher := utils.FilterTeacher(utils.ParseQueryParams(r, "teacher_id"))

	var teacher_id int
	if !filterTeacher {
		m, err := h.service.GetAll(limit, page, sort, field, teacher_id)
		if err != nil {
			handleLogger.Error(err)
			exceptions.BadQueryException(w, err)
			return
		}
		responses.GetAllResponse(w, m, limit, page)
		handleLogger.Info("Get All Materials Success")
		return
	}

	teacher_id, err = utils.ConvertToInt(utils.ParseQueryParams(r, "teacher_id"))
	if err != nil {
		err = errors.New("invalid teacher_id query param")
		handleLogger.Error(err)
		exceptions.BadQueryException(w, err)
		return
	}
	teaService := models.NewTeacherService(repository.NewTeacherRepository())
	exists, err := teaService.TeacherExistsByID(teacher_id)

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

	m, err := h.service.GetAll(limit, page, sort, field, teacher_id)

	if err != nil {
		handleLogger.Error(err)
		exceptions.BadQueryException(w, err)
		return
	}

	responses.GetAllResponse(w, m, limit, page)
	handleLogger.Info("Get All Materials Success")
}

func (h *MaterialHandlerImpl) GetMaterialByID(w http.ResponseWriter, r *http.Request) {
	materialID, err := utils.ConvertToInt(chi.URLParam(r, "id"))
	if err != nil {
		exceptions.BadParamException(w)
		log.Error("invalid ID")
		return
	}

	handleLogger := log.WithFields(log.Fields{"get": fmt.Sprintf("/materials/%v", materialID)})

	exists, err := h.service.MaterialExistsByID(materialID)

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

	m, err := h.service.GetByID(materialID)

	if err != nil {
		exceptions.NotFoundException(w, err)
		handleLogger.Error(err)
		return
	}

	responses.GetByIDResponse(w, m)
	handleLogger.Info("Get Material By ID success")
}

func (h *MaterialHandlerImpl) PostMaterial(w http.ResponseWriter, r *http.Request) {
	handleLogger := log.WithFields(log.Fields{"post": "/materials"})

	var updatedMaterial models.PayloadMaterial
	err := json.NewDecoder(r.Body).Decode(&updatedMaterial)

	if err != nil {
		exceptions.BadBodyException(w)
		handleLogger.Error(err)
		return
	}
	teaService := models.NewTeacherService(repository.NewTeacherRepository())

	exists, err := teaService.TeacherExistsByID(updatedMaterial.Teacher_id)

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

	mat := h.service.NewMaterial(updatedMaterial)

	if err != nil {
		exceptions.BadBodyException(w)
		handleLogger.Error(err)
		return
	}

	responses.CreatedResponse(w, mat)
}

func (h *MaterialHandlerImpl) UpdateMaterial(w http.ResponseWriter, r *http.Request) {
	materialID, err := utils.ConvertToInt(chi.URLParam(r, "id"))
	if err != nil {
		exceptions.BadParamException(w)
		log.Error("invalid ID")
		return
	}

	handleLogger := log.WithFields(log.Fields{"put": fmt.Sprintf("/materials/%v", materialID)})

	exists, err := h.service.MaterialExistsByID(materialID)

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

	teaService := models.NewTeacherService(repository.NewTeacherRepository())
	exists, err = teaService.TeacherExistsByID(updatedMaterial.Teacher_id)

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

	mat, err := h.service.UpdateMaterial(materialID, updatedMaterial)

	if err != nil {
		exceptions.BadRequestException(w, err)
		handleLogger.Error(err)
		return
	}

	responses.UpdateByIDResponse(w, mat)
	handleLogger.Info("Updated Material success")
}
