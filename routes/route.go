package routes

import (
	"log"
	"net/http"

	"github.com/DanielDanteDosSantosViana/darth_vader/db"
	"github.com/DanielDanteDosSantosViana/darth_vader/handlers"
	"github.com/DanielDanteDosSantosViana/darth_vader/models"
	"github.com/gorilla/mux"
)

func StartRoutes() {
	sessionW, err := db.NewSessionWriteDB()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar abrir conexão com o dbWrite . %v", err)
	}

	err = sessionW.Ping()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar verificar conexão dbWrite . %v", err)
	}

	sessionR, err := db.NewSessionReadDB()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar abrir conexão com o dbRead . %v", err)
	}

	err = sessionR.Ping()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar verificar conexão dbRead . %v", err)
	}
	fileModel := models.NewFileModel(sessionW, sessionR)

	//Inicia o mux
	r := mux.NewRouter()
	r.HandleFunc("/file", handlers.NewFileHandler(fileModel).List).Methods("GET")

	//Atribuí o mux ao http
	http.Handle("/", r)
}
