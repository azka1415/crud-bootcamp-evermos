package responses

import (
	"net/http"
)

func DeleteMaterialResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
