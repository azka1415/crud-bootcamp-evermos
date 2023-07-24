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

type TeacherHandler struct{}

func NewTeacherHandler() *TeacherHandler {
	return &TeacherHandler{}
}

func (t *TeacherHandler) GetTeacher(w http.ResponseWriter, r *http.Request) {

	handleLogger := log.WithFields(log.Fields{"get": "/teachers"})

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

	teaService := models.NewTeacherService()

	m, err := teaService.GetAll(limit, page, sort, field)

	if err != nil {
		handleLogger.Error(err)
		exceptions.BadQueryException(w, err)
		return
	}

	responses.GetAllResponse(w, m, limit, page)
	handleLogger.Info("Get All Teachers Success")
}

func (t *TeacherHandler) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID, err := utils.ConvertToInt(chi.URLParam(r, "id"))
	if err != nil {
		exceptions.BadParamException(w)
		log.Error("invalid ID")
		return
	}
	handleLogger := log.WithFields(log.Fields{"delete": fmt.Sprintf("/teachers/%v", teacherID)})
	teaService := models.NewTeacherService()

	exists, err := teaService.TeacherExistsByID(teacherID)

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
	_, err = teaService.DeleteTeacher(teacherID)

	if err != nil {
		exceptions.BadRequestException(w, err)
		handleLogger.Error(err)
		return
	}

	responses.NoContentResponse(w)
}

func (t *TeacherHandler) GetTeacherByID(w http.ResponseWriter, r *http.Request) {
	teacherID, err := utils.ConvertToInt(chi.URLParam(r, "id"))
	if err != nil {
		exceptions.BadParamException(w)
		log.Error("invalid ID")
		return
	}

	handleLogger := log.WithFields(log.Fields{"get": fmt.Sprintf("/teachers/%v", teacherID)})
	teacherService := models.NewTeacherService()

	exists, err := teacherService.TeacherExistsByID(teacherID)

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

	m, err := teacherService.GetByID(teacherID)

	if err != nil {
		exceptions.NotFoundException(w, err)
		handleLogger.Error(err)
		return
	}

	responses.GetByIDResponse(w, m)
	handleLogger.Info("Get Teacher By ID success")
}

func (t *TeacherHandler) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	teaID, err := utils.ConvertToInt(chi.URLParam(r, "id"))
	if err != nil {
		exceptions.BadParamException(w)
		log.Error("invalid ID")
		return
	}

	handleLogger := log.WithFields(log.Fields{"put": fmt.Sprintf("/teachers/%v", teaID)})

	teaService := models.NewTeacherService()

	exists, err := teaService.TeacherExistsByID(teaID)

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

	var updatedTeacher models.PayloadTeacher
	err = json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		exceptions.BadBodyException(w)
		handleLogger.Error(err)
		return
	}

	tea, err := teaService.UpdateTeacher(teaID, updatedTeacher)

	if err != nil {
		exceptions.BadRequestException(w, err)
		handleLogger.Error(err)
		return
	}

	responses.UpdateByIDResponse(w, tea)
	handleLogger.Info("Update Teacher success")
}

func (t *TeacherHandler) PostTeacher(w http.ResponseWriter, r *http.Request) {
	handleLogger := log.WithFields(log.Fields{"post": "/teachers"})

	var updatedTeacher models.PayloadTeacher
	err := json.NewDecoder(r.Body).Decode(&updatedTeacher)

	if err != nil {
		exceptions.BadBodyException(w)
		handleLogger.Error(err)
		return
	}
	teaService := models.NewTeacherService()

	tea, err := teaService.NewTeacher(updatedTeacher)

	if err != nil {
		exceptions.BadBodyException(w)
		handleLogger.Error(err)
		return
	}

	responses.CreatedResponse(w, tea)
	handleLogger.Info("Created Teacher success")
}
