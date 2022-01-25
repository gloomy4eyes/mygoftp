package cli

import (
	"github.com/spf13/cobra"
)

type Cli struct {
	root *cobra.Command
}

func NewCli() *Cli {
	c := new(Cli)

	root := rootCommand()

	return c
}

func (cli *Cli) Execute() error {
	return cli.root.Execute()
}
