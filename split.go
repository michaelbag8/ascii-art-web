package main

import (
	"strings"
)

func SplitInput(str string) []string {
	return strings.Split(str, "\\n")
}
