package models

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
	"path/filepath"
)

type File struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Nome      string        `json:"nome"`
	Directory string        `json:"directory"`
	Status    *StatusFile   `json:"status"`
	Bank      string        `json:"bank"`
	Sequencial string       `json:"sequencial"`
}

func NewFile(nome, directory string, status *StatusFile) *File {
	nomeTmp :=strings.TrimSuffix(nome,filepath.Ext(nome))
	info:= strings.Split(nomeTmp, "_")
	bank:=info[2]
	seq:=info[3]
	return &File{Nome: nome, Directory: directory, Status: status,Bank:bank,Sequencial:seq}
}
