package responses

import (
	"encoding/json"
	"net/http"
)

func GetByIDResponse(w http.ResponseWriter, m interface{}) {
	w.WriteHeader(http.StatusOK)
	res := Response{Message: "Success", Data: m}
	json.NewEncoder(w).Encode(res)
}
