package command

import (
	"strings"

	"github.com/jngreenwood/nazgul/internal"
	"github.com/mitchellh/cli"
)

type NodeWalkCommand struct {
	Ui cli.Ui
}

func (f *NodeWalkCommand) Help() string {
	helpText := `
	Usage: nazgul node poll <name>
	`
	return strings.TrimSpace(helpText)
}

func (f *NodeWalkCommand) Synopsis() string {
	return "Walks a named device"
}

func (f *NodeWalkCommand) Name() string { return "node walk" }

func (f *NodeWalkCommand) Run(args []string) int {

	agent := internal.NewAgent()
	agent.Start()

	if len(args) < 1 {
		f.Ui.Error("You must specify a node name")
		return 1
	}

	node, err := agent.GetNode(args[0])
	if err != nil {
		f.Ui.Error(err.Error())
		return 1
	}

	err = agent.Walk(node)
	if err != nil {
		f.Ui.Error(err.Error())
		return 1
	}

	return 0

}
