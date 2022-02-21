package imgconv

import (
	"strings"
)

func trimError(err error) string {
	s := err.Error()
	for i, c := range s {
		if c == ' ' {
			return s[i+1:]
		}
	}
	return s
}

func replaceSuffix(s, old, new string) string {
	return strings.TrimSuffix(s, old) + new
}
