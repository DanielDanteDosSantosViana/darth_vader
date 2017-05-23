package models

import "gopkg.in/mgo.v2/bson"

type File struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Nome      string        `json:"nome"`
	Directory string        `json:"directory"`
	Status    *StatusFile
}

func NewFile(nome, directory string, status *StatusFile) *File {
	return &File{Nome: nome, Directory: directory, Status: status}
}
