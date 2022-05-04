package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type NodeListCommand struct {
	Ui cli.Ui
}

func (f *NodeListCommand) Help() string {
	helpText := `
	Usage: nazgul node list [options] [args]
	`
	return strings.TrimSpace(helpText)
}

func (f *NodeListCommand) Synopsis() string {
	return "Lists nodes"
}

func (f *NodeListCommand) Name() string { return "node list" }

func (f *NodeListCommand) Run(args []string) int {
	return cli.RunResultHelp
}
