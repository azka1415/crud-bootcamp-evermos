package handlers

import (
	"errors"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/exceptions"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
	"github.com/azka1415/crud-bootcamp-evermos/tools/utils"
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
	sort := utils.ParseQueryParams(r, "sort")

	materialService := models.MaterialService{}
	m, err := materialService.GetAll(limit, offset)

}
