package rest

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetID extracts :id params from the request url
func GetID(r *http.Request) (int, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return 0, err
	}
	return id, nil
}
