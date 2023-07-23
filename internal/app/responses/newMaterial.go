package responses

import (
	"encoding/json"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
)

func NewMaterialResponse(w http.ResponseWriter, m models.Material) {
	w.WriteHeader(http.StatusCreated)
	res := Response{Message: "Success", Data: m}
	json.NewEncoder(w).Encode(res)
}
