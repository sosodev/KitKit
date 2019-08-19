package utilities

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCopyFile(t *testing.T) {
	tempDir, err := ioutil.TempDir("testdata", "kitkit")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourcePath := path.Join("testdata", "kitkit", "binaries", "testBinary-kktag:latest")
	targetPath := path.Join(tempDir, "testbinary")
	err = CopyFile(sourcePath, targetPath)
	if err != nil {
		t.Fatalf("failed to copy file: %s", err)
	}

	sourceFile, err := os.Stat(sourcePath)
	if err != nil {
		t.Fatalf("failed to stat source file: %s", err)
	}

	targetFile, err := os.Stat(targetPath)
	if err != nil {
		t.Fatalf("failed to stat copied file: %s", err)
	}

	if sourceFile.Size() != targetFile.Size() {
		t.Errorf("source and target file size mismatch: %s", err)
	}
}
