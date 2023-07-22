package responses

import (
	"encoding/json"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
)

type PaginatedResponse struct {
	Response
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func GetAllMaterialResponse(w http.ResponseWriter, m []models.Material, limit, page int) {
	w.WriteHeader(http.StatusOK)
	res := PaginatedResponse{
		Response{Message: "Success", Data: m}, limit, page,
	}
	json.NewEncoder(w).Encode(res)
}
