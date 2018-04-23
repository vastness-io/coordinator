package vcs_event

import (
	"strings"
)

func RemoveDirectoryPrefix(files []string) []string {

	var out []string

	for _, s := range files {
		index := strings.LastIndex(s, "/")

		if index != -1 {
			r := []rune(s)
			out = append(out, string(r[index+1:]))
			continue
		}

		out = append(out, s)
	}

	return out
}
