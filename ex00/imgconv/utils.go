package imgconv

import (
	"errors"
	"os"
	"strings"
)

func getFileExtentFromFile(filePath string) string {
	dot_at := 0
	for i := len(filePath) - 1; i > -1; i-- {
		if filePath[i] == '.' {
			dot_at = i
			break
		}
	}
	return filePath[dot_at:]
}

func validateArgs() error {
	if len(os.Args) == 1 {
		return errors.New("error: invalid argument")
	}
	return nil
}

func isValidFileExtent(path, ext string) bool {
	switch ext {
	case ".jpg", ".jpeg":
		return strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg")
	default:
		return strings.HasSuffix(path, ext)
	}
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
