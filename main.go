package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oscar-mgh/api-rest-golang-mysql/common"
	"github.com/oscar-mgh/api-rest-golang-mysql/routes"
)

var port = ":8080"

func main() {
	common.Migrate()
	common.Seed()
	router := mux.NewRouter()
	routes.SetPersonRoutes(router)

	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Println("Server started on port: ", port)
	log.Println(server.ListenAndServe())
}
