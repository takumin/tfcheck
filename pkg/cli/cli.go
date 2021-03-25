package cli

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/takumin/tfcheck/pkg/cli/completion"
	"github.com/takumin/tfcheck/pkg/cli/modules"
	"github.com/takumin/tfcheck/pkg/cli/providers"
	"github.com/takumin/tfcheck/pkg/config"
)

type App interface {
	Run(args []string) error
	RunContext(ctx context.Context, args []string) error
}

type app struct {
	cli *cli.App
}

func NewApp(name, usage, version, revision string, c *config.Config) App {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "working-directory",
			Aliases:     []string{"dir", "d"},
			Usage:       "working directory",
			Value:       c.WorkingDirectory,
			Destination: &c.WorkingDirectory,
			EnvVars:     []string{"WORKING_DIR", "DIR"},
		},
	}

	cmds := []*cli.Command{}
	cmds = append(cmds, completion.NewCommands(c, flags)...)
	cmds = append(cmds, providers.NewCommands(c, flags)...)
	cmds = append(cmds, modules.NewCommands(c, flags)...)

	return &app{
		cli: &cli.App{
			Name:                 name,
			Usage:                usage,
			Version:              fmt.Sprintf("%s (%s)", version, revision),
			Flags:                flags,
			Commands:             cmds,
			EnableBashCompletion: true,
		},
	}
}

func (a *app) Run(args []string) error {
	return a.cli.Run(args)
}

func (a *app) RunContext(ctx context.Context, args []string) error {
	return a.cli.RunContext(ctx, args)
}
