package cmd

import (
	"github.com/gocolly/colly/v2"
	"github.com/urfave/cli/v2"
)

var crawlCmd = &cli.Command{
	Name:  "log",
	Flags: []cli.Flag{},
	Action: func(ctx *cli.Context) error {
		_ = colly.NewCollector()
		return nil
	},
}
