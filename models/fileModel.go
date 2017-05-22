package models

import (
	"log"

	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionName = "files"
)

type FileModel struct {
	writeDB *mgo.Session
	readDB  *mgo.Session
}

func NewFileModel(writeDB *mgo.Session, readDB *mgo.Session) *FileModel {
	return &FileModel{writeDB, readDB}
}

func (fm *FileModel) Create(file *File) error {
	if err := fm.writeDB.DB(config.Conf.Db.Name).C(collectionName).Insert(file); err != nil {
		return err
	}
	return nil
}

func (fm *FileModel) List(directory string) ([]File, error) {
	var files []File
	if err := fm.readDB.DB(config.Conf.Db.Name).C(collectionName).Find(bson.M{"directory": directory}).All(&files); err != nil {
		return nil, err
	}
	log.Print(files)

	return files, nil
}

func (fm *FileModel) GetBy(name string) ([]File, error) {
	var files []File
	if err := fm.readDB.DB(config.Conf.Db.Name).C(collectionName).Find(bson.M{"name": name}).All(&files); err != nil {
		return files, err
	}
	return files, nil
}
