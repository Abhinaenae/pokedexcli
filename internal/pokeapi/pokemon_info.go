package pokeapi

type PokemonInfo struct {
	BaseExperience int    `json:"base_experience"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Species        struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
}
