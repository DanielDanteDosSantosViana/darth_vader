package models

import (
	"database/sql"
	"errors"
	"log"
)

type AreaModel struct {
	readDB  *sql.DB
	writeDB *sql.DB
	area    *Area
}

type Area struct {
	Id        int    `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
}

func NewAreaModel(read *sql.DB, write *sql.DB) *AreaModel {
	return &AreaModel{read, write, &Area{}}
}

func (a *AreaModel) SetArea(area *Area) *Area {
	a.area = area
	return a.area
}

func (a *AreaModel) Criar() (*Area, error) {
	var stmt *sql.Stmt
	stmt, err := a.writeDB.Prepare(`INSERT INTO area(descricao) VALUES(?)`)
	if err != nil {
		log.Printf("Error na sequence. %v", err)
		return &Area{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(a.area.Descricao)
	if err != nil {
		log.Printf("dbWrite error ao inserir. %v", err)
		return &Area{}, errors.New("dbWrite error ao inserir")
	}
	return a.area, nil
}

func (a *AreaModel) Atualizar() (*Area, error) {
	var stmt *sql.Stmt
	stmt, err := a.writeDB.Prepare(`UPDATE area set descricao=? where id = ?`)
	if err != nil {
		log.Printf("Error na sequence. %v", err)
		return &Area{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(a.area.Descricao, a.area.Id)
	if err != nil {
		log.Printf("dbWrite error ao inserir. %v", err)
		return &Area{}, errors.New("dbWrite error ao inserir")
	}
	return a.area, nil
}

func (a *AreaModel) Deletar() error {
	var stmt *sql.Stmt
	stmt, err := a.writeDB.Prepare(`DELETE FROM area where id = ?`)
	if err != nil {
		log.Printf("Error na sequence. %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(a.area.Id)
	if err != nil {
		log.Printf("dbWrite error ao deletar. %v", err)
		return errors.New("dbWrite error ao deletar")
	}

	return nil
}

func (a *AreaModel) Listar() ([]Area, error) {
	var areas []Area

	return areas, nil
}
