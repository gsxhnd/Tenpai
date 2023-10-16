package main

import (
	"github.com/gsxhnd/tenpai/tenhou/src/catalog"
)

const INPUT_PATH = "input/html"

func main() {
	c := catalog.NewCatalog(INPUT_PATH)
	c.Archive()
}
