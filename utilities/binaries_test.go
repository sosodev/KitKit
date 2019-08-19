package utilities

import (
	"os"
	"path"
	"testing"
)

func TestGetBinaries(t *testing.T) {
	oldHome := os.Getenv("KITKIT_HOME")
	defer os.Setenv("KITKIT_HOME", oldHome)
	err := os.Setenv("KITKIT_HOME", path.Join("testdata/kitkit"))
	if err != nil {
		t.Fatalf("failed to setup kitkit home: %s", err)
	}

	binaries, err := GetBinaries()
	if err != nil {
		t.Fatalf("failed to get binaries: %s", err)
	}

	if binaries[0].Name() != "testBinary-kktag:latest" {
		t.Errorf("incorrect binary name wanted: %s got: %s", "testBinary-kktag:latest", binaries[0].Name())
	}
}

func TestSplitTrackedName(t *testing.T) {
	testCases := []struct {
		Have     string
		WantName string
		WantTag  string
	}{
		{
			Have:     "kikit-kktag:latest",
			WantName: "kikit",
			WantTag:  "latest",
		},
		{
			Have:     "kitkit-latest-darwin64-kktag:latest",
			WantName: "kitkit-latest-darwin64",
			WantTag:  "latest",
		},
	}

	for _, tc := range testCases {
		t.Run("split tracked name", func(t *testing.T) {
			name, tag := SplitTrackedName(tc.Have)

			if tc.WantName != name {
				t.Errorf("incorrect name want: %s got: %s", tc.WantName, name)
			}

			if tc.WantTag != tag {
				t.Errorf("incorrect tag want: %s got: %s", tc.WantTag, tag)
			}
		})
	}
}
