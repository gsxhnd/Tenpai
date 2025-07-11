package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "tenhou",
		Usage: "tenhou data",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("Hello friend!", time.Now())
			time.Sleep(1 * time.Second)
			fmt.Println("Hello friend!", time.Now())
			time.Sleep(1 * time.Second)
			fmt.Println("Hello friend!", time.Now())
			time.Sleep(1 * time.Second)
			fmt.Println("Hello friend!", time.Now())

			dlT, ok := ctx.Deadline()
			fmt.Println(dlT, ok)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
