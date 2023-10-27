package cmd

import (
	"github.com/gsxhnd/tenpai/tenhou/src/service"
	"github.com/urfave/cli/v2"
)

const DEST_PATH = "input/catalog"

var catalogCmd = &cli.Command{
	Name:  "log",
	Flags: []cli.Flag{},
	Action: func(ctx *cli.Context) error {
		c := service.NewCatalog(DEST_PATH)
		c.Archive()
		return nil
	},
}
