package api

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	resp := response{
		Status: http.StatusOK,
		Data:   "API v1",
	}
	resp.ServerJSON(w)
}
