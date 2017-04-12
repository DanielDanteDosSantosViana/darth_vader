package routes

import (
	"net/http"

	"github.com/DanielDanteDosSantosViana/darth_vader/handlers"
	"github.com/gorilla/mux"
)

func StartRoutes() {
	//Inicia o mux
	r := mux.NewRouter()
	r.HandleFunc("/file", handlers.NewFileHandler().List).Methods("GET")
	r.HandleFunc("/file/{name}", handlers.NewFileHandler().GetByName).Methods("GET")

	//Atribuí o mux ao http
	http.Handle("/", r)
}
