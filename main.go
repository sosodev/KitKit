package main

import (
	"log"
	"os"

	"github.com/simplyserenity/kitkit/config"

	"github.com/mitchellh/cli"
)

func main() {
	os.Exit(realMain(os.Args[1:]))
}

func realMain(args []string) int {
	c := cli.NewCLI("KitKit", "1.1")
	c.Args = args
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	err := config.KitkitSetup()
	if err != nil {
		log.Fatalf("Failed to setup the $KITKIT_HOME directory: %s", err)
		return 127
	}

	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{
				Ui: ui,
			}, nil
		},
		"list": func() (cli.Command, error) {
			return &ListCommand{
				Ui: ui,
			}, nil
		},
		"set": func() (cli.Command, error) {
			return &SetCommand{
				Ui: ui,
			}, nil
		},
		"remove": func() (cli.Command, error) {
			return &RemoveCommand{
				Ui: ui,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	return exitStatus
}
