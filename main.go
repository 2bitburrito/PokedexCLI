package main

import (
	"fmt"
	"time"

	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
	"github.com/2bitburrito/pokedex-cli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	fmt.Println("cache created")
	newClient := pokeapi.NewClient(5 * time.Second, cache)
	cfg := &pokeapi.Config{
		PokeClient: newClient,
	}
	startRepl(cfg)
}
