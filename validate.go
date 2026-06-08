package main

import (
	"fmt"
	"strings"
)

func ValidateInput(input string, banner string) error {
	input = strings.ReplaceAll(input, "\r", "") //this

	if input == "" {
		return fmt.Errorf("input is empty")
	}
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		return fmt.Errorf("not a banner method")
	}
	for _, ch := range input {
		if ch == '\n' || ch == '\r'{ //ch == '\r' this
			continue
		}
		if ch < 32 || ch > 126 {
			return fmt.Errorf("non ascii character: %c", ch)
		}
	}
	return nil
}
