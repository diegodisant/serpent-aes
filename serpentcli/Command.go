package serpentcli

import "github.com/urfave/cli"

type Command struct {
	Name      string
	Aliases   []string
	UsageText string
	Exec      func(context *cli.Context) error
}
