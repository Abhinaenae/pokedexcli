package main

import "github.com/abhinaenae/pokedexcli/internal/pokeapi"

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	runDex(&cfg)
}
