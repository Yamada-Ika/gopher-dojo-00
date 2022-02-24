package imgconv_bonus

import (
	"strings"
)

func isValidFileExtent(path, ext string) bool {
	if ext == ".jpg" || ext == ".jpeg" {
		return strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg")
	}
	return strings.HasSuffix(path, ext)
}

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
