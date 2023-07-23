package exceptions

import (
	"encoding/json"
	"net/http"
)

func BadQueryException(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	ex := Exception{Message: err.Error(), Error: "Bad Request"}
	json.NewEncoder(w).Encode(ex)
}
