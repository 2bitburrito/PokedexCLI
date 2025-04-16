package pokeapi

import (
	"net/http"
	"time"

	"github.com/2bitburrito/pokedex-cli/internal/pokecache"
)

type Client struct {
	HttpClient http.Client
	Cache         *pokecache.Cache
}

func NewClient(timeout time.Duration, cache *pokecache.Cache) *Client {
	return &Client{
		HttpClient: http.Client{
			Timeout: timeout,
		},
		Cache: cache,
	}
}
