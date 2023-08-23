package routes

import (
	"github.com/gorilla/mux"
	"github.com/oscar-mgh/api-rest-golang-mysql/controller"
)

func SetPersonRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api/person").Subrouter()
	subRoute.HandleFunc("", controller.Save).Methods("POST")
	subRoute.HandleFunc("", controller.GetAll).Methods("GET")
	subRoute.HandleFunc("/{id}", controller.Get).Methods("GET")
	subRoute.HandleFunc("/{id}", controller.Update).Methods("PUT")
	subRoute.HandleFunc("/{id}", controller.Delete).Methods("DELETE")
}
