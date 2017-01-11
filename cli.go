package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Converts gior files to other formats."
	app.Commands = []cli.Command{
		ConvertCommand,
	}
	app.Run(os.Args)
}
