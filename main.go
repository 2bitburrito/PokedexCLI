package main

import (
	"time"

	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
	"github.com/2bitburrito/pokedex-cli/internal/types"
)

func main() {
	newClient := pokeapi.NewClient(5 * time.Second)
	cfg := types.Config{
		PokeClient: newClient,
	}
	startRepl(&cfg)
}
