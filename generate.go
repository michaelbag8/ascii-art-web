package main

import (
	"fmt"
	"os"
	"strings"
)

func Generate(input string, banner map[rune][]string) string {
	var output strings.Builder
	segments := SplitInput(input)

	for i, segment := range segments {
		if segment == "" {

			if len(segments) == 1 {
				return ""
			}

			if i < len(segments)-1 {
				output.WriteString("\n")

			}
			if len(segments) == 2 && segments[0] == "" {
				output.WriteString("\n")
				return ""

			}

			if i == len(segments)-1 {
				output.WriteString("\n")
			}

			continue
		}

		for _, ch := range segment {
			if _, ok := banner[ch]; !ok {
				fmt.Fprintf(os.Stderr, "Non ASCII character %q\n", ch)
				os.Exit(1)
			}
		}

		rows := renderLines(segment, banner)
		for _, row := range rows {
			output.WriteString(row + "\n")
		}
	}
	return output.String()
}
