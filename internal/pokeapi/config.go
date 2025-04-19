package pokeapi

type Config struct {
	PokeClient       *Client
	Pokedex          PokedexInterface
	LocationsNext    *string
	LocationsPrev    *string
	ExploredLocation *string
	CatchingPokemon  *string
	PokemonToInspect *string
}

type PokedexInterface interface {
	AddPokemon(pokedata CatchResponse)
	HasPokemon(name string) bool
	ListPokemon()
	Inspect(pokemon string)
}
