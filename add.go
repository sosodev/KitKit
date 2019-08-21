package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/simplyserenity/kitkit/utilities"

	"github.com/simplyserenity/kitkit/config"

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
	var set bool

	// create the flagset
	f := flag.NewFlagSet("AddFlags", flag.ContinueOnError)
	f.StringVar(&name, "name", "", "name")
	f.StringVar(&tag, "tag", "latest", "tag")
	f.BoolVar(&set, "set-now", false, "set-now")
	f.Usage = func() { c.Ui.Error(c.Help()) }

	// separate the flags from the normal args so ordering doesn't matter
	args, flags := utilities.SeparateFlags(args)
	if len(args) == 0 {
		c.Ui.Error(c.Help())
		return 1
	}

	// parse the flags
	if err := f.Parse(flags); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	binaryPath := args[0]
	// check if binary exists
	binary, err := os.Stat(binaryPath)
	if err != nil {
		c.Ui.Error("The specified binary doesn't exist")
		return 1
	}

	// take its name if none specified
	if name == "" {
		name = binary.Name()
	}

	if !utilities.ValidIdentifier(name) {
		c.Ui.Error(fmt.Sprintf("invalid name \"%s\" - valid names can contain alphanumeric characters, hyphens, and periods", name))
		return 1
	}

	if !utilities.ValidIdentifier(tag) {
		c.Ui.Error(fmt.Sprintf("invalid tag \"%s\" - valid tags can contain alphanumeric characters, hyphens, and periods", tag))
		return 1
	}

	// create the tagged name and path
	taggedName := name + "-kktag:" + tag
	taggedPath := path.Join(config.BinariesPath(), taggedName)

	// copy it to the binaries folder under the tagged name
	err = utilities.CopyFile(binaryPath, taggedPath)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to copy the binary: %s", err.Error()))
	}

	c.Ui.Output(fmt.Sprintf("Added %s tagged as %s", name, tag))

	if set {
		sc := &SetCommand{
			Ui: c.Ui,
		}

		return sc.Run([]string{name, tag})
	}

	return 0
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
	
	--set-now 	Immediately runs "kitkit set" after adding the specified binary
`
}

func (c *AddCommand) Synopsis() string {
	return "Adds a binary to be tracked by kitkit"
}
