package main

import (
	"time"

	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
	"github.com/2bitburrito/pokedex-cli/internal/pokecache"
	"github.com/2bitburrito/pokedex-cli/pokedex"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	newClient := pokeapi.NewClient(5*time.Second, cache)
	pokedex := pokedex.NewPokedex()
	cfg := &pokeapi.Config{
		PokeClient: newClient,
		Pokedex:    pokedex,
	}
	startRepl(cfg)
}
