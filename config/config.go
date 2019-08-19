package config

import (
	"os"
	"os/user"
	"path"
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

func BinariesPath() string {
	return path.Join(KitkitHome(), "binaries")
}

func BinPath() string {
	return path.Join(KitkitHome(), "bin")
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
