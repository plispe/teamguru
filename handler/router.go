package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler() http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	// Handlers
	r.HandleFunc("/", mainHandler).Methods(http.MethodGet)
	r.HandleFunc("/healtz", healthcheckHandler).Methods(http.MethodGet)
	r.HandleFunc("/version", versionHandler).Methods(http.MethodGet)

	return r
}
