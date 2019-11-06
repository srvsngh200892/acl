package router

import (
	"github.com/gorilla/mux"
	"github.com/srvsngh200892/acl/src/handler"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.HomepageHandler)
	router.HandleFunc("/roles", handler.CreateRoles).Methods("POST")
	router.HandleFunc("/users", handler.CreateUsers).Methods("POST")
	router.HandleFunc("/subordinates/{id}", handler.ListSubOrdinates).Methods("GET")
	router.HandleFunc("/hc/acl", handler.HealthCheckHandler).Methods("GET")

	return router
}