package commands

import (
	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
)

func CommandInspect(cfg *pokeapi.Config) error {
	pokemon := cfg.PokemonToInspect
	if pokemon == nil {
		return nil
	}
	cfg.Pokedex.Inspect(*pokemon)
	return nil
}
