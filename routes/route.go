package routes

import (
	"net/http"

	"github.com/DanielDanteDosSantosViana/darth_vader/handlers"
	"github.com/DanielDanteDosSantosViana/darth_vader/models"
	"github.com/gorilla/mux"
)

func StartRoutes() {
	fileModel := models.NewFileModel()

	//Inicia o mux
	r := mux.NewRouter()
	r.HandleFunc("/file", handlers.NewFileHandler(fileModel).List).Methods("GET")

	//Atribuí o mux ao http
	http.Handle("/", r)
}
