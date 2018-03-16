package event

import "strings"

func RemoveDirectoryPrefix(s string) string {

	index := strings.LastIndex(s, "/")

	if index != -1 {
		r := []rune(s)
		return string(r[index+1:])
	}

	return s
}
