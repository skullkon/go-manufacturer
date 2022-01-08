package api

import (
	"github.com/gorilla/mux"
	v1 "github.com/skullkon/go-manufacturer/internal/api/v1"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/v1/manufacturer", v1.Handler)

	return router
}
