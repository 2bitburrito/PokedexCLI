package types

import "github.com/2bitburrito/pokedex-cli/internal/pokeapi"

type Config struct {
	PokeClient    Client
	LocationsNext string
	LocationsPrev string
}

type Client interface {
	GetLocations(url string) (pokeapi.LocationsResponse, error)
}
