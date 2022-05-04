package command

import (
	"strings"

	"github.com/jngreenwood/nazgul/internal"
	"github.com/mitchellh/cli"
)

type SNMPModelTestCommand struct {
	Ui cli.Ui
}

func (f *SNMPModelTestCommand) Help() string {
	helpText := `
	Usage: nazgul model test [options] [args]
	`
	return strings.TrimSpace(helpText)
}

func (f *SNMPModelTestCommand) Synopsis() string {
	return "Tests a snmp model"
}

func (f *SNMPModelTestCommand) Name() string { return "snmp_test" }

func (f *SNMPModelTestCommand) Run(args []string) int {

	_, err := internal.LoadNodeModel("models/default.hcl")
	if err != nil {
		f.Ui.Error(err.Error())
		return 1
	}
	f.Ui.Output("Model passed!")

	return 0
}
