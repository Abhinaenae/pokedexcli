package main

import (
	"time"

	"github.com/abhinaenae/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	runDex(&cfg)
}
