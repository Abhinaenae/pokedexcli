package main

import "github.com/abhinaenae/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	History             []string
}
