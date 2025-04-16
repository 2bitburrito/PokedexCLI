package pokeapi

type Config struct {
	PokeClient       *Client
	LocationsNext    *string
	LocationsPrev    *string
	ExploredLocation *string
	CatchingPokemon  *string
}
