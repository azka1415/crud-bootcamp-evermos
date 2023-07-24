package handlers

import (
	"errors"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/exceptions"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/responses"
	"github.com/azka1415/crud-bootcamp-evermos/tools/utils"
	"github.com/azka1415/crud-bootcamp-evermos/tools/utils/enums/sort"
	log "github.com/sirupsen/logrus"
)

type TeacherHandler struct{}

func NewTeacherHandler() *TeacherHandler {
	return &TeacherHandler{}
}

func (t *TeacherHandler) GetTeacher(w http.ResponseWriter, r *http.Request) {

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

	teaService := models.NewTeacherService()

	m, err := teaService.GetAll(limit, page, sort, field)

	if err != nil {
		handleLogger.Error(err)
		exceptions.BadQueryException(w, err)
		return
	}

	responses.GetAllResponse(w, m, limit, page)
	handleLogger.Info("Get All Materials Success")
}

func (t *TeacherHandler) DeleteTeacher(w http.ResponseWriter, r *http.Request) {

}
