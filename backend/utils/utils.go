package utils

import (
	"net/http"
)

const (
	OneMb = 1048576
	PORT  = "80"
)

func SetRequestHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}
