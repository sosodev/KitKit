package utilities

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/simplyserenity/kitkit/config"
)

// GetBinaries returns os.Fileinfo for all of the binaries found in KITKIT_HOME/binaries
func GetBinaries() ([]os.FileInfo, error) {
	binariesPath := config.BinariesPath()
	binaries, err := ioutil.ReadDir(binariesPath)
	if err != nil {
		return nil, err
	}
	return binaries, nil
}

// Splits a tracked name into it's respective name and tag
func SplitTrackedName(trackedName string) (string, string) {
	parts := strings.Split(trackedName, "-kktag:")
	return parts[0], parts[1]
}
