package providers

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/takumin/tfcheck/pkg/config"
	"github.com/takumin/tfcheck/pkg/fswalk"
	"github.com/takumin/tfcheck/pkg/tfcheck"
)

func NewCommands(c *config.Config, f []cli.Flag) []*cli.Command {
	flags := []cli.Flag{}
	return []*cli.Command{
		{
			Name:    "providers",
			Aliases: []string{"p"},
			Usage:   "Check version terraform providers",
			Flags:   append(flags, f...),
			Action:  action(c),
		},
	}
}

func action(c *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		dirs, err := fswalk.FsWalk(c.WorkingDirectory, c.ExcludeDirectories)
		if err != nil {
			return fmt.Errorf("Failed to fswalk.FsWalk(): %w", err)
		}

		parser := tfcheck.NewParser(dirs)
		versions, err := parser.Parse()
		if err != nil {
			return fmt.Errorf("Failed to tfcheck.Parse(): %w", err)
		}

		providers := make(map[string]string)

		for _, version := range versions {
			for k, v := range version.Providers {
				providers[k] = v.Source
			}
		}

		for k := range providers {
			fmt.Println(k)
		}

		return nil
	}
}
