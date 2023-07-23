package responses

import (
	"encoding/json"
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
)

func UpdateMaterialResponse(w http.ResponseWriter, m models.Material) {
	w.WriteHeader(http.StatusOK)
	res := Response{Message: "Success", Data: m}
	json.NewEncoder(w).Encode(res)
}
