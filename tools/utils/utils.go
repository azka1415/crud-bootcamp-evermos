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

func CheckFieldQuery(s string) string {
	switch strings.ToLower(s) {
	case "id":
		return s
	case "title":
		return s
	case "created_at":
		return s
	case "updated_at":
		return s
	case "teacher_id":
		return s
	default:
		return "id"
	}
}
