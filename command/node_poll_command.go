package command

import (
	"strings"

	"github.com/jngreenwood/nazgul/internal"
	"github.com/mitchellh/cli"
)

type NodePollCommand struct {
	Ui cli.Ui
}

func (f *NodePollCommand) Help() string {
	helpText := `
	Usage: nazgul node poll <name>
	`
	return strings.TrimSpace(helpText)
}

func (f *NodePollCommand) Synopsis() string {
	return "Polls a named device"
}

func (f *NodePollCommand) Name() string { return "node poll" }

func (f *NodePollCommand) Run(args []string) int {

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

	err = agent.Poll(node)
	if err != nil {
		f.Ui.Error(err.Error())
		return 1
	}

	return 0

}
