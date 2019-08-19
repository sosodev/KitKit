package config

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func testSetKitkitHome(t *testing.T) func() error {
	oldHome := os.Getenv("KITKIT_HOME")
	err := os.Setenv("KITKIT_HOME", path.Join("testdata", "kitkit"))
	if err != nil {
		t.Fatal(err)
	}

	return func() error {
		return os.Setenv("KITKIT_HOME", oldHome)
	}
}

func TestKitkitHome(t *testing.T) {
	defer testSetKitkitHome(t)()

	home := KitkitHome()
	wantHome := path.Join("testdata", "kitkit")
	if home != wantHome {
		t.Errorf("incorrect kitkit home got: %s want: %s", home, wantHome)
	}
}

func TestBinPath(t *testing.T) {
	defer testSetKitkitHome(t)()

	binPath := BinPath()
	wantPath := path.Join(KitkitHome(), "bin")
	if binPath != wantPath {
		t.Errorf("incorrect binpath got: %s want: %s", binPath, wantPath)
	}
}

func TestBinariesPath(t *testing.T) {
	defer testSetKitkitHome(t)()

	binariesPath := BinariesPath()
	wantPath := path.Join(KitkitHome(), "binaries")
	if binariesPath != wantPath {
		t.Errorf("incorrect binaries path got: %s want: %s", binariesPath, wantPath)
	}
}

func TestKitkitSetup(t *testing.T) {
	defer testSetKitkitHome(t)()

	// should do nothing, or at least not fail here
	err := KitkitSetup()
	if err != nil {
		t.Error(err)
	}

	tempDir, err := ioutil.TempDir("testdata", "kitkit")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tempDir)

	err = os.Setenv("KITKIT_HOME", tempDir)
	if err != nil {
		t.Fatal(err)
	}

	err = KitkitSetup()
	if err != nil {
		t.Fatalf("failed to setup: %s", err)
	}

	_, err = os.Stat(BinPath())
	if err != nil {
		t.Error("couldn't find bin path")
	}

	_, err = os.Stat(BinariesPath())
	if err != nil {
		t.Error("couldn't find binaries path")
	}
}
