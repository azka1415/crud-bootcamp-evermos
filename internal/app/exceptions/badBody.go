package exceptions

import (
	"encoding/json"
	"net/http"
)

func BadBodyException(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	ex := Exception{Message: "Invalid Request Body", Error: "Bad Request"}
	json.NewEncoder(w).Encode(ex)
}
