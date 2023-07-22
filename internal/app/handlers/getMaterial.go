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

func HandleGetMaterial(w http.ResponseWriter, r *http.Request) {

	handleLogger := log.WithFields(log.Fields{"get": "/materials"})

	page, err := utils.ConvertToInt(utils.ParseQueryParams(r, "page"))
	if err != nil {
		handleLogger.Error(errors.New("please check your query"))
		exceptions.BadQueryException(w)
		return
	}

	limit, err := utils.ConvertToInt(utils.ParseQueryParams(r, "limit"))
	if err != nil {
		handleLogger.Error(errors.New("please check your query"))
		exceptions.BadQueryException(w)
		return
	}

	sort := sort.GetSortDirection(utils.ParseQueryParams(r, "sort"))
	field := utils.ParseQueryParams(r, "field")

	materialService := models.MaterialService{}

	m, err := materialService.GetAll(limit, page, sort, field)

	if err != nil {
		handleLogger.Info(err)
		handleLogger.Error(errors.New("please check your query"))
		exceptions.BadQueryException(w)
		return
	}

	responses.GetAllMaterialResponse(w, m, limit, page)
	handleLogger.Info("Get All Materials Success")
}
