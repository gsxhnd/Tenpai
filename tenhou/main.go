package main

import (
	"os"

	"github.com/gsxhnd/tenpai/tenhou/src/cmd"
)

func main() {
	cmd.RootCmd.Run(os.Args)
}
