package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Demo() {
	app := &cli.App{
		Name:  "TCPScan",
		Usage: "let network discover easy.",
		Action: func(c *cli.Context) error {
			fmt.Println(c.Args().Len())
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
