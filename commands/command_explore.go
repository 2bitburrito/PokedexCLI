package commands

import (
	"fmt"

	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
)

func CommandExplore(cfg *pokeapi.Config)error {
	exploreResp, err := cfg.PokeClient.ExploreLocation(cfg.ExploredLocation)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring: %s\n",*cfg.ExploredLocation)
	fmt.Println("Found Pokemon:")
	for _, encounter := range exploreResp.PokemonEncounters{
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
 
}