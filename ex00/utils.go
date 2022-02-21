package main

import (
	"strings"
)

func replaceSuffix(s, old, new string) string {
	return strings.TrimSuffix(s, old) + new
}
