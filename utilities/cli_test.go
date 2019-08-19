package utilities

import (
	"reflect"
	"testing"
)

func TestSeparateFlags(t *testing.T) {
	// cases
	// three types of flags: -flag=something, -flag something, and -flag(boolean)
	// and more with -- but should be handled the same
	// can be at any position in the array (ideally?)
	testCases := []struct {
		Name      string
		HaveArgs  []string
		WantArgs  []string
		WantFlags []string
	}{
		{
			Name:      "Before flags",
			HaveArgs:  []string{"-string1=x", "--bool2", "--string2=y", "--bool1", "test", "args"},
			WantArgs:  []string{"test", "args"},
			WantFlags: []string{"-string1=x", "--bool2", "--string2=y", "--bool1"},
		},
		{
			Name:      "After flags",
			HaveArgs:  []string{"test", "args", "-string1=x", "--bool2", "--string2=y", "--bool1"},
			WantArgs:  []string{"test", "args"},
			WantFlags: []string{"-string1=x", "--bool2", "--string2=y", "--bool1"},
		},
		{
			Name:      "Sporadic flags",
			HaveArgs:  []string{"test", "-string1=x", "--bool2", "args", "--string2=y", "--bool1"},
			WantArgs:  []string{"test", "args"},
			WantFlags: []string{"-string1=x", "--bool2", "--string2=y", "--bool1"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var gotFlags []string
			tc.HaveArgs, gotFlags = SeparateFlags(tc.HaveArgs)
			if !reflect.DeepEqual(gotFlags, tc.WantFlags) {
				t.Errorf("flags incorrect: want %v got %v", tc.WantFlags, gotFlags)
			}

			if !reflect.DeepEqual(tc.HaveArgs, tc.WantArgs) {
				t.Errorf("args incorrect: want %v got %v", tc.WantArgs, tc.HaveArgs)
			}
		})
	}
}
