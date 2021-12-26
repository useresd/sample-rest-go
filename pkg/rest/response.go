package rest

import (
	"encoding/json"
	"net/http"
)

func WriteJson(rw http.ResponseWriter, statusCode int, content interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(content)
}
