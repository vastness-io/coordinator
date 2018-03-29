package vcs_event

import (
	"strings"
)

func RemoveDirectoryPrefix(s string) string {

	index := strings.LastIndex(s, "/")

	if index != -1 {
		r := []rune(s)
		return string(r[index+1:])
	}

	return s
}

func RemoveDuplicates(files []string, f string) []string {
	ok, index := containsFile(f, files)

	if ok {
		files = append(files[:index], files[index+1:]...)
	} else {
		files = append(files, f)
	}

	return files
}

func containsFile(f string, files []string) (bool, int) {
	for i, _ := range files {
		if files[i] == f {
			return true, i
		}
	}
	return false, -1
}
