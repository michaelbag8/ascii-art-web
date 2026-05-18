package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file\n")
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("invalid content")
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	content = strings.TrimLeft(content, "\n")

	rawFile := strings.Split(content, "\n\n")

	blockMaps := make(map[rune][]string)
	charCode := ' '

	for _, raw := range rawFile {
		lines := strings.Split(raw, "\n")
		if len(lines) < 8 {
			return nil, fmt.Errorf("invalid content")
		}
		blockMaps[charCode] = lines[:8]
		charCode++
	}
	return blockMaps, nil
}
