package main

import "strings"

func Generate(input string, banner map[rune][]string) string {
	var output strings.Builder

	segments := SplitInput(input)

	for i, segment := range segments {
		if segment == "" {
			if i == 0 || segments[i-1] != "" {
				output.WriteString("\n")
			}
			continue
		}

		rows := renderLines(segment, banner)
		for _, row := range rows {
			output.WriteString(row)
			output.WriteString("\n")
		}
	}
	return output.String()
}
