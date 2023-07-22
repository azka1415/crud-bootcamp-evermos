package utils

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var errorLogger = log.WithFields(log.Fields{"Error": "Log"})

func ParseQueryParams(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func ConvertToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
