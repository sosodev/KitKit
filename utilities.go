package main

import (
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"strings"
)

// KitkitHome returns the path the kitkit directory
func KitkitHome() string {
	home := os.Getenv("KITKIT_HOME")
	if home == "" {
		usr, err := user.Current()
		if err != nil {
			panic(err) // I can't even find out why this would error
		}
		home = path.Join(usr.HomeDir, ".kitkit")
	}
	return home
}

// KitkitSetup ensures that $KITKIT_HOME and all of it's subdirectories exist
func KitkitSetup() error {
	home := KitkitHome()
	binariesPath := path.Join(home, "binaries")
	binPath := path.Join(home, "bin")

	_, err := os.Stat(binariesPath)
	if err != nil {
		err = os.MkdirAll(binariesPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	_, err = os.Stat(binPath)
	if err != nil {
		err = os.MkdirAll(binPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetBinaries returns os.Fileinfo for all of the binaries found in KITKIT_HOME/binaries
func GetBinaries() ([]os.FileInfo, error) {
	home := KitkitHome()
	binariesPath := path.Join(home, "binaries")
	binaries, err := ioutil.ReadDir(binariesPath)
	if err != nil {
		return nil, err
	}
	return binaries, nil
}

// SeparateFlags splits a full set of arguments into a two string slices of just args and the flags
func SeparateFlags(args []string) (justArgs []string, flags []string) {
	for _, arg := range args {
		if arg[0] == '-' {
			flags = append(flags, arg)
		} else {
			justArgs = append(justArgs, arg)
		}
	}

	return
}

// CopyFile copies a file
func CopyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = os.Chmod(destination, os.ModePerm)
	if err != nil {
		return err
	}

	return destinationFile.Close()
}

// Splits a tracked name into it's respective name and tag
func SplitTrackedName(trackedName string) (string, string) {
	parts := strings.Split(trackedName, "-kktag:")
	return parts[0], parts[1]
}
