package cmd

import "github.com/urfave/cli/v2"

var RootCmd = &cli.App{
	Name: "tenhou",
	Commands: []*cli.Command{
		crawlCmd,
		catalogCmd,
		logCmd,
	},
}
