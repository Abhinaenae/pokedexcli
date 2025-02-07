package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	runDex()

}

func runDex() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}
	for {
		fmt.Printf("Pokedex > ")

		//Wait for user input
		if !scanner.Scan() {
			break
		}

		userInput := scanner.Text()

		cleanedInput := strings.Fields(strings.ToLower(strings.TrimSpace(userInput)))
		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]

		if cmd, exists := commands[command]; exists {
			if err := cmd.callback(cfg); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
