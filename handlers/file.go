package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DanielDanteDosSantosViana/darth_vader/models"
	"github.com/DanielDanteDosSantosViana/darth_vader/util"
)

type File struct {
	fileModel *models.FileModel
}

func NewFileHandler(fileModel *models.FileModel) *File {
	return &File{fileModel}
}

func (f *File) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	directory := r.URL.Query().Get("directory")
	if directory == "" {
		files, err := f.fileModel.ListAll()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if len(files) == 0 {
			util.ResponseNotFound(w, "Don't have files")
			return
		}
		filesJ, _ := json.Marshal(files)
		util.ResponseOK(w, filesJ)
	}else{
		files, err := f.fileModel.List(directory)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if len(files) == 0 {
			util.ResponseNotFound(w, "Don't have files")
			return
		}
		filesJ, _ := json.Marshal(files)
		util.ResponseOK(w, filesJ)
	}

}
