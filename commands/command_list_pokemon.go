package commands

import "github.com/2bitburrito/pokedex-cli/internal/pokeapi"

func CommandListPokemon(cfg *pokeapi.Config) error {
	cfg.Pokedex.ListPokemon()
	return nil
}
