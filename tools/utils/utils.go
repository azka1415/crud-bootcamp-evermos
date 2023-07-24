package utils

import (
	"net/http"
	"strconv"
	"strings"
)

func ParseQueryParams(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func ConvertToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func FilterTeacher(s string) bool {
	return strings.ToLower(s) != ""
}
