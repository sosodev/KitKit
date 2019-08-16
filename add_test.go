package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/mitchellh/cli"
)

// kitkit add [relative path to binary file] --name(-n)=[binary's default name] --tag(-t)=[string for the tag(version)]
// the binary's name is what kitkit will list it as and add it to the path as
// the tag is how the separate versions of the binary will be tracked (similar to docker tags)
func TestAddCommand_Run(t *testing.T) {
	testDir, cleanup := testKitKitDirectory(t)
	defer cleanup()

	ui := new(cli.MockUi)
	c := &AddCommand{
		Ui: ui,
	}

	args := []string{
		"testdata/testBinary",
		"-name=test-binary",
		"-tag=1.0",
	}

	if returnCode := c.Run(args); returnCode != 0 {
		t.Fatalf("bad return code: %d\n\n%s", returnCode, ui.ErrorWriter.String())
	}

	want := "binaries/test-binary-kktag:1.0"
	// check if the file was copied and tagged
	if _, err := os.Stat(path.Join(testDir, want)); err != nil {
		t.Fatalf("couldn't find the copied binary file: %s", want)
	}
}

// creates the testing directory
// sets KITKIT_HOME to the testing directory
// returns the testdir and a function that will unset the env and delete the testing directory
func testKitKitDirectory(t *testing.T) (string, func()) {
	oldHome := os.Getenv("KITKIT_HOME")
	dir, err := ioutil.TempDir("testdata", "kitkit")
	if err != nil {
		t.Fatalf("failed to create tempdir: %s", err)
	}

	err = os.Setenv("KITKIT_HOME", dir)
	if err != nil {
		t.Fatal(err)
	}

	if err = KitkitSetup(); err != nil {
		t.Fatal(err)
	}

	return dir, func() {
		err = os.Setenv("KITKIT_HOME", oldHome)
		if err != nil {
			t.Fatal(err)
		}

		err = os.RemoveAll(dir)
		if err != nil {
			t.Fatalf("failed to remove the testdata/kitkit directory: %s", err)
		}
	}
}
