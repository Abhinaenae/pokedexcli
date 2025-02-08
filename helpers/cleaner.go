package helper

import (
	"fmt"
	"strings"
)

func CleanInput(text string) []string {
	if len(text) == 0 {
		fmt.Println("Warning: Recieved empty string")
		return []string{}
	}

	//Change all letters to lowercase and converts the string to type []string
	cleaned := strings.TrimSpace(text)
	lowered := strings.ToLower(cleaned)
	cleanOutput := strings.Fields(lowered)

	if len(cleanOutput) == 0 {
		fmt.Println("Warning: Input contained only punctuation and was removed.")
		return []string{}
	}

	return cleanOutput
}
