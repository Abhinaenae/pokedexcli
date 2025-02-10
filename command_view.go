package main

import (
	"fmt"
)

func commandView(cfg *config, args ...string) error {
	fmt.Println("Pokemon registered in Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Species.Name)
	}

	return nil
}
