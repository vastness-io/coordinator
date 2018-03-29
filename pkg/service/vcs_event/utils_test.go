package vcs_event

import (
	"reflect"
	"testing"
)

func TestExcludeRemovedFilesFromLangRequest(t *testing.T) {
	tests := []struct {
		in          []string
		toBeRemoved string
		expected    []string
	}{
		{
			in: []string{
				"main.clj",
				"a.py",
				"pom.xml",
			},
			toBeRemoved: "pom.xml",
			expected: []string{
				"main.clj",
				"a.py",
			},
		},
	}

	for _, test := range tests {

		result := RemoveDuplicates(test.in, test.toBeRemoved)
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Expected %v, got %v", test.expected, result)
		}
	}

}
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
