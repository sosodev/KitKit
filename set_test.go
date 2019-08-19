package main

import (
	"os"
	"path"
	"testing"

	"github.com/simplyserenity/kitkit/config"

	"github.com/mitchellh/cli"
)

func TestSetCommand_Run(t *testing.T) {
	testDir, cleanup := testKitKitDirectory(t)
	defer cleanup()

	ui := new(cli.MockUi)
	c := &SetCommand{
		Ui: ui,
	}

	testAddToBinaries(t, testDir)

	args := []string{
		"testbinary",
		"latest",
	}

	if returnCode := c.Run(args); returnCode != 0 {
		t.Fatalf("bad return code: %d\n\n%s", returnCode, ui.ErrorWriter.String())
	}

	_, err := os.Stat(path.Join(config.BinPath(), "testbinary"))
	if err != nil {
		t.Errorf("failed to find testbinary in binpath: %s", err)
	}
}
