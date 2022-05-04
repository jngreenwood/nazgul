package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type NodeCommand struct {
	Ui cli.Ui
}

func (f *NodeCommand) Help() string {
	helpText := `
	Usage: nazgul node <subcommand> [options] [args]

	Subcommands:
		poll
		list
	`
	return strings.TrimSpace(helpText)
}

func (f *NodeCommand) Synopsis() string {
	return "ToDO"
}

func (f *NodeCommand) Name() string { return "node" }

func (f *NodeCommand) Run(args []string) int {
	return cli.RunResultHelp
}
