package exceptions

import (
	"encoding/json"
	"net/http"
)

func NotFoundException(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	ex := Exception{Message: err.Error(), Error: "Not Found"}
	json.NewEncoder(w).Encode(ex)
}
