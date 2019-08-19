package main

import (
	"os"
	"path"
	"testing"

	"github.com/simplyserenity/kitkit/utilities"

	"github.com/mitchellh/cli"
)

func TestRemoveCommand_Run(t *testing.T) {
	testDir, cleanup := testKitKitDirectory(t)
	defer cleanup()

	ui := new(cli.MockUi)
	c := &RemoveCommand{
		Ui: ui,
	}

	binaryPath := testAddToBinaries(t, testDir)
	args := []string{
		"testbinary",
		"latest",
	}

	returnCode := c.Run(args)
	if returnCode != 0 {
		t.Fatalf("bad return code: %d\n\n%s", returnCode, ui.ErrorWriter.String())
	}

	_, err := os.Stat(binaryPath)
	if err == nil {
		t.Error("binary still exists")
	}
}

func testAddToBinaries(t *testing.T, testDir string) string {
	binaryPath := path.Join(testDir, "binaries", "testbinary-kktag:latest")
	err := utilities.CopyFile(path.Join("testdata/testBinary"), binaryPath)
	if err != nil {
		t.Fatalf("failed to copy binary: %s", err)
	}
	return binaryPath
}
