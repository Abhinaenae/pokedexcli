package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Why are you throwing a pokeball at nothing?")
	}
	pokemonName := args[0]

	/*
		pokemonNameResp, err := cfg.pokeapiClient.GetLocationArea(pokemonName)
		if err != nil {
			return err
		}
	*/

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	fmt.Printf("Caught %s!!\n", pokemonName)

	return nil
}
