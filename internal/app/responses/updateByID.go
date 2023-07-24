package responses

import (
	"encoding/json"
	"net/http"
)

func UpdateMaterialResponse(w http.ResponseWriter, m interface{}) {
	w.WriteHeader(http.StatusOK)
	res := Response{Message: "Success", Data: m}
	json.NewEncoder(w).Encode(res)
}
