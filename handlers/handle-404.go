package handlers

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var Handle404 = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	// rw.Write([]byte("Page not found, my friend"))
	json.NewEncoder(rw).Encode(HTTPError{
		Code:    http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
	})
})
