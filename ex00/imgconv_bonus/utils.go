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

func replaceFileExtent(filePath, oldExt, newExt string) string {
	if strings.HasSuffix(filePath, ".jpg") && oldExt == ".jpeg" {
		return replaceSuffix(filePath, ".jpg", newExt)
	} else if strings.HasSuffix(filePath, ".jpeg") && oldExt == ".jpg" {
		return replaceSuffix(filePath, ".jpeg", newExt)
	}
	return replaceSuffix(filePath, oldExt, newExt)
}

func replaceSuffix(s, old, new string) string {
	return strings.TrimSuffix(s, old) + new
}
