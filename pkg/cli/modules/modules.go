package modules

import (
	"fmt"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/takumin/tfcheck/pkg/config"
	"github.com/takumin/tfcheck/pkg/fswalk"
	"github.com/takumin/tfcheck/pkg/tfcheck"
)

func NewCommands(c *config.Config, f []cli.Flag) []*cli.Command {
	flags := []cli.Flag{}
	return []*cli.Command{
		{
			Name:    "modules",
			Aliases: []string{"m"},
			Usage:   "Check version terraform modules",
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

		modules := make(map[string]string)
		for _, version := range versions {
			for _, v := range version.Modules {
				if strings.Contains(v.Source, "//") {
					modules[v.Source[:strings.Index(v.Source, "//")]] = v.Version
				} else {
					modules[v.Source] = v.Version
				}
			}
		}

		results := make([]string, 0, len(modules))
		for k := range modules {
			results = append(results, k)
		}

		sort.Strings(results)
		for _, v := range results {
			fmt.Println(v)
		}

		return nil
	}
}
