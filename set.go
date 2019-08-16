package main

import (
	"fmt"
	"path"

	"github.com/mitchellh/cli"
)

type SetCommand struct {
	Ui cli.Ui
}

func (c *SetCommand) Run(args []string) int {
	home := KitkitHome()
	binPath := path.Join(home, "bin")
	binariesPath := path.Join(home, "binaries")

	targetName := args[0]
	targetTag := args[1]

	binaries, err := GetBinaries()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to load binaries in $KITKIT_HOME/binaries: %s", err))
		return 127
	}

	for _, binary := range binaries {
		name, tag := SplitTrackedName(binary.Name())

		if name == targetName && tag == targetTag {
			sourcePath := path.Join(binariesPath, binary.Name())
			destinationPath := path.Join(binPath, name)
			err := CopyFile(sourcePath, destinationPath)
			if err != nil {
				c.Ui.Error(fmt.Sprintf("Failed to copy the specified binary: %s", err))
				return 127
			}
			return 0
		}
	}

	c.Ui.Error("The specified binary could not be found.")
	return 1
}

func (c *SetCommand) Help() string {
	return `
Usage: kitkit set [binary-name] [tag]

	Puts the tagged binary into $KITKIT_HOME/bin.
	$KITKIT_HOME/bin should be on your path for it to be available.
`
}

func (c *SetCommand) Synopsis() string {
	return "Sets a binary on the path"
}
