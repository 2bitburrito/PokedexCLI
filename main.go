package main

import (
	"time"

	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
)

func main() {
	newClient := pokeapi.NewClient(5 * time.Second)
	cfg := &pokeapi.Config{
		PokeClient: newClient,
	}
	startRepl(cfg)
}
