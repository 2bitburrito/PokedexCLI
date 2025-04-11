package pokeapi

type Config struct {
	PokeClient    *Client
	LocationsNext *string
	LocationsPrev *string
}

type Clients interface {
	GetLocations() (LocationsResponse, error)
}
