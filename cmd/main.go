package main

import (
	"fmt"
	"os"

	"github.com/858chain/erc20-transfer/utils"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ERC20 token transfer service"
	app.Version = Version
	app.Commands = []cli.Command{
		startCmd,
	}

	app.Flags = []cli.Flag{
		logLevelFlag,
		logDirFlag,
		envFlag,
	}

	app.Before = func(c *cli.Context) error {
		return utils.InitLogger(c.String("log-dir"), c.String("log-level"), "json")
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
