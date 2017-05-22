package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DanielDanteDosSantosViana/darth_vader/models"
	"github.com/gorilla/mux"
)

type File struct {
	fileModel *models.FileModel
}

func NewFileHandler(fileModel *models.FileModel) *File {
	return &File{fileModel}
}

func (f *File) List(w http.ResponseWriter, r *http.Request) {
	directory := r.URL.Query().Get("directory")
	if directory == "" {
		responseBadRequest(w, "Not found directory")
		return
	}

	files, err := f.fileModel.List(directory)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	filesJ, _ := json.Marshal(files)
	responseOK(w, filesJ)
}
func (file *File) GetByName(w http.ResponseWriter, r *http.Request) {

}

func getUrlParameter(r *http.Request, parameter string) string {
	vars := mux.Vars(r)
	return vars[parameter]
}

func responseOK(w http.ResponseWriter, a ...interface{}) {
	if a != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", a)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func responseBadRequest(w http.ResponseWriter, a ...interface{}) {
	if a != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", a)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
}
