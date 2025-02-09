package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\n\nUSAGE:\n\n")
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}
