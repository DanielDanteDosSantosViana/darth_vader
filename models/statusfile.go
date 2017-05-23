package models

import "time"

type StatusFile struct {
	Total        string `toml:"total"`
	Value        string `toml:"value"`
	DataReceiver string `toml:"dataReceiver"`
}

func NewStatusFile(total, value string) *StatusFile {
	dateNow := time.Now()
	dataFormat := dateNow.Format("02-01-2006")
	return &StatusFile{total, value, dataFormat}
}
