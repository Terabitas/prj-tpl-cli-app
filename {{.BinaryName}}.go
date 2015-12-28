package main

import (
	"os"

	"github.com/codegangsta/cli"
	"{{.OrgPath}}/{{.ProjectName}}/version"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "{{.BinaryName}}"
	app.Usage = "{{.BinaryDescription}}"
	app.Version = version.Version
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "verbosity",
			Value: 0,
			Usage: "logging level",
		},
	}
	app.Before = func(c *cli.Context) error {
		// setup logging here
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "sample",
			Aliases: []string{"s"},
			Usage:   "sample command",
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatalf("Could not run {{.BinaryName}} cmd, %s", err)
	}
}
