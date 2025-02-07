package main

import (
	"fmt"
	"regexp"
	"strings"
)

func cleanInput(text string) []string {
	if len(text) == 0 {
		fmt.Println("Warning: Recieved empty string")
		return []string{}
	}
	// Define a regex pattern to remove non-word characters except spaces
	re := regexp.MustCompile(`[^\w\s]`)

	// Replace all matches (punctuation) with an empty string
	cleaned := re.ReplaceAllString(text, "")

	// Split the string into words using spaces
	words := strings.Fields(cleaned)
	if len(words) == 0 {
		fmt.Println("Warning: Input contained only punctuation and was removed.")
		return []string{}
	}

	return words
}
