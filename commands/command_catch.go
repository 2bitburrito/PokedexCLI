package commands

import (
	"fmt"
	"math/rand/v2"

	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
)

func CommandCatch(cfg *pokeapi.Config) error {
	pokemonResp, err := cfg.PokeClient.CatchPokemon(cfg.CatchingPokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", *cfg.CatchingPokemon)
	chance := rand.IntN(300)
	fmt.Print(chance, pokemonResp.BaseExperience, "\n")
	if chance > pokemonResp.BaseExperience {
		fmt.Printf("%s caught!\n", *cfg.CatchingPokemon)
		cfg.Pokedex.AddPokemon(pokemonResp)
	} else {
		fmt.Printf("%s escaped.\n", *cfg.CatchingPokemon)
	}
	return nil

}
