package vcs_event

import (
	"reflect"
	"testing"
)

func TestRemoveDirectoryPrefix(t *testing.T) {
	tests := []struct {
		in  []string
		out []string
	}{
		{
			in:  []string{"i have no parent"},
			out: []string{"i have no parent"},
		},
		{
			in:  []string{"parent/pom.xml"},
			out: []string{"pom.xml"},
		},
		{
			[]string{"1/2/3/4/more/depth/pom.xml"},
			[]string{"pom.xml"},
		},
		{
			in:  []string{""},
			out: []string{""},
		},
		{
			in:  []string{"x/"},
			out: []string{""},
		},
		{
			in:  []string{"//"},
			out: []string{""},
		},
		{
			in:  []string{"pom.xml"},
			out: []string{"pom.xml"},
		},
	}

	for _, test := range tests {

		result := RemoveDirectoryPrefix(test.in)
		if !reflect.DeepEqual(test.out, result) {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}

}
