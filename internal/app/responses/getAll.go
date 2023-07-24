package responses

import (
	"encoding/json"
	"net/http"
)

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type PaginatedResponse struct {
	Response
	Pagination
}

func GetAllResponse(w http.ResponseWriter, m interface{}, limit, page int) {
	w.WriteHeader(http.StatusOK)
	res := PaginatedResponse{
		Response{Message: "Success", Data: m}, Pagination{Limit: limit, Page: page},
	}
	json.NewEncoder(w).Encode(res)
}
