package models

import "gopkg.in/mgo.v2/bson"

type File struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Nome   string        `json:"nome"`
	Evento string        `json:"evento"`
	Data   string        `json:"data"`
}

func NewFile() *File {
	return &File{}
}
