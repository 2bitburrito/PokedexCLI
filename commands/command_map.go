package commands

import (
	"fmt"

	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
)

func CommandMap(cfg *pokeapi.Config) error {
	locations, err := cfg.PokeClient.GetLocations(cfg.LocationsNext)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.LocationsNext = &locations.Next
	cfg.LocationsPrev = &locations.Previous

	return nil
}

func CommandMapBack(cfg *pokeapi.Config) error {
	locations, err := cfg.PokeClient.GetLocations(cfg.LocationsPrev)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.LocationsNext = &locations.Next
	cfg.LocationsPrev = &locations.Previous

	return nil
}
