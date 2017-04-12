package handlers

import (
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
	name := getUrlParameter(r, "name")
	files, err := f.fileModel.GetBy(name)
	if err != nil {
		log.Println(err.Error())
		return
	}

}
func (file *File) GetByName(w http.ResponseWriter, r *http.Request) {

}

func getUrlParameter(r *http.Request, parameter string) string {
	vars := mux.Vars(r)
	return vars[parameter]
}
