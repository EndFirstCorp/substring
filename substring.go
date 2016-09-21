package substring

import (
	"strings"
)

func After(in string, search string) string {
	startIndex := strings.Index(in, search)
	if startIndex == -1 {
		return ""
	}
	startIndex += len(search)
	return in[startIndex:len(in)]
}

func Between(source, from, to string) string {
	startIndex := strings.Index(source, from)
	if startIndex == -1 {
		return ""
	}
	fromIndex := startIndex + len(from)
	endIndex := strings.Index(source[fromIndex:], to)
	if endIndex == -1 {
		return ""
	}
	toIndex := endIndex + fromIndex
	return source[fromIndex:toIndex]
}
