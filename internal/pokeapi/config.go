package pokeapi

type Config struct {
	PokeClient       *Client
	Pokedex          PokedexInterface
	LocationsNext    *string
	LocationsPrev    *string
	ExploredLocation *string
	CatchingPokemon  *string
}

type PokedexInterface interface {
	AddPokemon(pokedata CatchResponse)
	HasPokemon(name string) bool
	ListPokemon()
}
