package models

import (
	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/DanielDanteDosSantosViana/darth_vader/db"
	"log"
)

const (
	collectionName = "files"
)

type FileModel struct {
	writeDB *mgo.Session
	readDB  *mgo.Session
}

func NewFileModel() *FileModel {

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
	return &FileModel{sessionW, sessionR}
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
	return files, nil
}

func (fm *FileModel) ListAll() ([]File, error) {
	var files []File
	if err := fm.readDB.DB(config.Conf.Db.Name).C(collectionName).Find(nil).All(&files); err != nil {
		return nil, err
	}
	return files, nil
}
