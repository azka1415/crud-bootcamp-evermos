package responses

import (
	"net/http"
)

func NoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
