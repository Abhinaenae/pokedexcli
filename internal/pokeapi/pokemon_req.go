package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonInfo, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonInfo{}, fmt.Errorf("bad status code: %v", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokemonInfo := PokemonInfo{}
	if err := json.Unmarshal(data, &pokemonInfo); err != nil {
		return PokemonInfo{}, err
	}

	return pokemonInfo, nil
}
