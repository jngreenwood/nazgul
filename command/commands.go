package command

import "github.com/mitchellh/cli"

func Commands(ui cli.Ui) map[string]cli.CommandFactory {

	all := map[string]cli.CommandFactory{
		"node": func() (cli.Command, error) {
			return &NodeCommand{
				Ui: ui,
			}, nil
		},
		"node poll": func() (cli.Command, error) {
			return &NodePollCommand{
				Ui: ui,
			}, nil
		},
		"node walk": func() (cli.Command, error) {
			return &NodeWalkCommand{
				Ui: ui,
			}, nil
		},
		"model test": func() (cli.Command, error) {
			return &SNMPModelTestCommand{
				Ui: ui,
			}, nil
		},
	}
	return all
}
