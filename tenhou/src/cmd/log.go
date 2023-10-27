package cmd

import (
	"github.com/gsxhnd/tenpai/tenhou/src/di"
	"github.com/urfave/cli/v2"
)

var logCmd = &cli.Command{
	Name:  "log",
	Flags: []cli.Flag{},
	Action: func(ctx *cli.Context) error {
		di.InitLog()
		return nil
	},
}
