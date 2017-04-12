package db

import (
	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	"gopkg.in/mgo.v2"
)

func NewSessionWriteDB() (*mgo.Session, error) {
	db, err := mgo.Dial(config.Conf.Db.MongoWrite)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewSessionReadDB() (*mgo.Session, error) {
	db, err := mgo.Dial(config.Conf.Db.MongoRead)
	if err != nil {
		return nil, err
	}
	return db, nil
}
