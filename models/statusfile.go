package models

import "time"

type StatusFile struct {
	Total        string `toml:"total"`
	Value        string `toml:"value"`
	FileName     string `toml:"filename"`
	DataReceiver string `toml:"dataReceiver"`
}

func NewStatusFile(total, value, fileName string) *StatusFile {
	dateNow := time.Now()
	dataFormat := dateNow.Format("02-01-2006")
	return &StatusFile{total, value, fileName, dataFormat}
}
