package main

import (
	"bufio"
	"fmt"
	"os"

	helper "github.com/abhinaenae/pokedexcli/helpers"
)

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

		cleanedInput := helper.CleanInput(userInput)
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
