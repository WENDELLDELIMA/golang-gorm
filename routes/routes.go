package routes

import (
	"log"
	"net/http"

	"github.com/WENDELLDELIMA/go-rest-api/controllers"
	"github.com/WENDELLDELIMA/go-rest-api/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest(){
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/personalidades/{id}", controllers.RetornaUmapersonalidades).Methods("GET")
	r.HandleFunc("/personalidades", controllers.CriaNovaPersonalidade).Methods("POST")
	r.HandleFunc("/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("DELETE")
	r.HandleFunc("/personalidades/{id}", controllers.EditaPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}