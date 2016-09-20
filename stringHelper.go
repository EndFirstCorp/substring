package helpers

import (
	"strings"
)

func SubstringAfter(in string, search string) string {
	startIndex := strings.Index(in, search)
	if startIndex == -1 {
		return in
	}
	startIndex += len(search)
	return in[startIndex:len(in)]
}

func SubstringBetween(source, from, to string) string {
	fromIndex := strings.Index(source, from) + len(from)
	toIndex := strings.Index(source[fromIndex:], to) + fromIndex
	return source[fromIndex:toIndex]
}
