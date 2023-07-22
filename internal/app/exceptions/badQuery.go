package exceptions

import (
	"encoding/json"
	"net/http"
)

func BadQueryException(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	ex := Exception{Message: "Please check your query", Error: "Bad Request"}
	json.NewEncoder(w).Encode(ex)
}
