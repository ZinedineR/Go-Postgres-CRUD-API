package router

import (
	"go-postgres-crud/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/tv", controller.GetTVAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tv/{id}", controller.GetTV).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tv", controller.NewTV).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tv/{id}", controller.UpdateTVNew).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/tv/{id}", controller.RemoveTV).Methods("DELETE", "OPTIONS")

	return router
}
