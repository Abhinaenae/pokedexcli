package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("why are you throwing a pokeball at nothing?")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("throwing a pokeball at %s...\n", pokemon.Name)

	const threshold = 50
	captureRateLimit := rand.Intn(pokemon.BaseExperience)
	if captureRateLimit > threshold {
		fmt.Printf("%s broke out!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.caughtPokemon[pokemonName] = pokemon

	return nil
}
