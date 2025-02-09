package main

import (
	"bufio"
	"fmt"
	"os"

	helper "github.com/abhinaenae/pokedexcli/helpers"
)

func runDex(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")

		//Wait for user input
		if !scanner.Scan() {
			break
		}

		userInput := scanner.Text()

		cleanedInput := helper.CleanInput(userInput)
		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]

		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		if cmd, exists := commands[command]; exists {
			if err := cmd.callback(cfg, args...); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
