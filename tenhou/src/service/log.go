package service

import (
	"os"
	"path"

	"github.com/go-gota/gota/dataframe"
)

type Log struct{}

func NewLog() *Log {
	return &Log{}
}

func (l *Log) ReadCSV() {
	east4csv, err := os.OpenFile(path.Join("", "east4.csv"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	_ = dataframe.ReadCSV(east4csv)
}
