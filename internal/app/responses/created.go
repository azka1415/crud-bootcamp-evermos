package responses

import (
	"encoding/json"
	"net/http"
)

func CreatedResponse(w http.ResponseWriter, m interface{}) {
	w.WriteHeader(http.StatusCreated)
	res := Response{Message: "Success", Data: m}
	json.NewEncoder(w).Encode(res)
}
