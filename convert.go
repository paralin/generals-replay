package main

import (
	"encoding/json"
	"errors"
	"fmt"
	rengine "github.com/andyleap/giorengine"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

var ConvertArgs struct {
	InputFile  string
	OutputFile string
}

var ConvertCommand = cli.Command{
	Name:   "tojson",
	Usage:  "Converts a gior to json.",
	Action: runConvert,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "input",
			Destination: &ConvertArgs.InputFile,
			Usage:       "Read the gior file at `PATH`.",
		},
		cli.StringFlag{
			Name:        "output",
			Destination: &ConvertArgs.OutputFile,
			Usage:       "Write the json file at `PATH`.",
		},
	},
}

func runConvert(c *cli.Context) error {
	if ConvertArgs.InputFile == "" {
		return errors.New("Input file and output file must be given.")
	}

	inputFile, err := os.Open(ConvertArgs.InputFile)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	match, err := rengine.ParseReplay(inputFile)
	if err != nil {
		return err
	}

	dat, err := json.MarshalIndent(match, "", "\t")
	if err != nil {
		return err
	}

	if ConvertArgs.OutputFile == "" {
		fmt.Printf("%s\n", string(dat))
		return nil
	} else {
		return ioutil.WriteFile(ConvertArgs.OutputFile, dat, 0644)
	}
}
