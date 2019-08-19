package main

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/cli"
	"github.com/simplyserenity/kitkit/config"
	"github.com/simplyserenity/kitkit/utilities"
)

// RemoveCommands removes a binary from the list
// of binaries that are tracked by kikit
type RemoveCommand struct {
	Ui cli.Ui
}

func (c *RemoveCommand) Run(args []string) int {
	if len(args) < 2 {
		c.Ui.Error(c.Help())
		return 1
	}

	binaryName := args[0]
	binaryTag := args[1]

	binaries, err := utilities.GetBinaries()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("failed to get binaries: %s", err))
		return 127
	}

	for _, binary := range binaries {
		name, tag := utilities.SplitTrackedName(binary.Name())

		if name == binaryName && tag == binaryTag {
			if err := os.Remove(path.Join(config.BinariesPath(), binary.Name())); err != nil {
				c.Ui.Error(fmt.Sprintf("failed to delete binary: %s", err))
			}
			c.Ui.Output(fmt.Sprintf("Removed %s %s", name, tag))
			return 0
		}
	}

	c.Ui.Error("The specified binary was not found")
	return 1
}

func (c *RemoveCommand) Help() string {
	return `
Usage: kitkit remove [binary-name] [binary-tag]

	Removes the binary that matches the specified tag and name.
	This means the binary will be deleted from the binaries folder.
`
}

func (c *RemoveCommand) Synopsis() string {
	return "Removes a binary from the the list of tracked binaries"
}
