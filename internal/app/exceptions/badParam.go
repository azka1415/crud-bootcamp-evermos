package exceptions

import (
	"encoding/json"
	"net/http"
)

func BadParamException(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	ex := Exception{Message: "Invalid Param", Error: "Bad Request"}
	json.NewEncoder(w).Encode(ex)
}
