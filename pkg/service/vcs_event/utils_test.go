package vcs_event

import "testing"

func TestRemoveDirectoryPrefix(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "i have no parent",
			out: "i have no parent",
		},
		{
			in:  "parent/pom.xml",
			out: "pom.xml",
		},
		{
			"1/2/3/4/more/depth/pom.xml",
			"pom.xml",
		},
		{
			in:  "",
			out: "",
		},
		{
			in:  "x/",
			out: "",
		},
		{
			in:  "//",
			out: "",
		},
		{
			in:  "pom.xml",
			out: "pom.xml",
		},
	}

	for _, test := range tests {

		result := RemoveDirectoryPrefix(test.in)
		if test.out != result {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}

}
