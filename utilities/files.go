package utilities

import (
	"io"
	"os"
)

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
