package main

import (
	"log"
	"os"

	"github.com/takumin/tfcheck/pkg/cli"
	"github.com/takumin/tfcheck/pkg/config"
)

var (
	AppName  string = "tfcheck"
	Usage    string = "terraform dependency checker"
	Version  string = "unknown"
	Revision string = "unknown"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	app := cli.NewApp(
		AppName,
		Usage,
		Version,
		Revision,
		config.NewConfig(
			config.WorkingDirectory(dir),
			config.ExcludeDirectory(".git"),
			config.ExcludeDirectory(".terraform"),
		),
	)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
