package main

import (
	"os"
	"path"

	"github.com/go-gota/gota/dataframe"
)

const DEST_PATH = "input/catalog"

func main() {
	east4csv, err := os.OpenFile(path.Join(DEST_PATH, "east4.csv"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	_ = dataframe.ReadCSV(east4csv)
}
