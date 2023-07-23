package responses

import (
	"encoding/json"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
)

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type PaginatedResponse struct {
	Response
	Pagination
}

func GetAllMaterialResponse(w http.ResponseWriter, m []models.Material, limit, page int) {
	w.WriteHeader(http.StatusOK)
	res := PaginatedResponse{
		Response{Message: "Success", Data: m}, Pagination{Limit: limit, Page: page},
	}
	json.NewEncoder(w).Encode(res)
}
