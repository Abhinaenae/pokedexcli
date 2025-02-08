package pokeapi

import (
	"net/http"
	"time"

	"github.com/abhinaenae/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpclient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpclient: http.Client{
			Timeout: time.Minute,
		},
	}
}
