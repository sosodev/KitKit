package main

import (
	"fmt"

	"github.com/simplyserenity/kitkit/utilities"

	"github.com/mitchellh/cli"
)

// ListCommand lists all of the currently tracked binaries
type ListCommand struct {
	Ui cli.Ui
}

func (c *ListCommand) Run(args []string) int {
	binaries, err := utilities.GetBinaries()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to load binaries in $KITKIT_HOME/binaries: %s", err))
		return 127
	}

	c.Ui.Output(fmt.Sprintf("%-10s%-10s", "name", "tag"))
	for _, binary := range binaries {
		name, tag := utilities.SplitTrackedName(binary.Name())
		c.Ui.Output(fmt.Sprintf("%-10s%-10s", name, tag))
	}

	return 0
}

func (c *ListCommand) Help() string {
	return `
Usage: kitkit list

	Lists all of the binaries currently tracked by kitkit and their tags.
	Those binaries can be swapped onto the path using the set command.
`
}

func (c *ListCommand) Synopsis() string {
	return "Lists all currently tracked binaries"
}
