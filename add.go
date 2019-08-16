package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/cli"
)

// AddCommand adds a binary to be tracked by kitkit
// tracked binaries are stored in $KITKIT_HOME/binaries
// active binaries are stored in $KITKIT_HOME/bin (only one of each name at a time)
type AddCommand struct {
	Ui cli.Ui
}

func (c *AddCommand) Run(args []string) int {
	// flags
	var name, tag string

	// create the flagset
	f := flag.NewFlagSet("AddFlags", flag.ContinueOnError)
	f.StringVar(&name, "name", "", "name")
	f.StringVar(&tag, "tag", "latest", "tag")
	f.Usage = func() { c.Ui.Error(c.Help()) }

	// separate the flags from the normal args so ordering doesn't matter
	args, flags := SeparateFlags(args)
	if len(args) == 0 {
		return 1
	}

	// parse the flags
	if err := f.Parse(flags); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	binaryPath := args[0]

	binaryName, err := loadBinaryName(binaryPath)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to open the specified binary: %s", err.Error()))
		return 127
	}

	if name == "" {
		name = binaryName
	}

	taggedName := name + "-kktag:" + tag
	taggedBinaryPath := path.Join(KitkitHome(), "binaries", taggedName)

	err = CopyFile(binaryPath, taggedBinaryPath)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to copy specified binary: %s", err.Error()))
		return 127
	}

	return 0
}

func loadBinaryName(binaryPath string) (string, error) {
	binary, err := os.Open(binaryPath)
	if err != nil {
		return "", err
	}

	return binary.Name(), nil
}

func (c *AddCommand) Help() string {
	return `
Usage: kitkit add [path-to-binary] 

	Tags a binary to tracked by kitkit and stores a copy of it 
	in "$KITKIT_HOME/binaries".
	Tracked binaries can be listed with the list command.
	Tracked binaries can be swapped onto the path using the set command.

Options:

	-name=name 	The name kitkit will track the binary with. Defaults 
				to the current name of the binary.

	-tag=tag	The tag given to the binary to differentiate it from 
				others of the same name. Defaults to "latest". 
				Kind of like a docker tag.
`
}

func (c *AddCommand) Synopsis() string {
	return "Adds a binary to be tracked by kitkit"
}
