package pokedex

import (
	"fmt"
	"sync"

	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
)

type Pokedex struct {
	Entries map[string]Pokemon
	M       sync.Mutex
}

type Pokemon struct {
	Name       string
	Experience int
	Stats      []struct {
		BaseStat int
		Effort   int
		Stat     struct {
			Name string
			URL  string
		}
	}
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		Entries: make(map[string]Pokemon),
		M:       sync.Mutex{},
	}
}

func (p *Pokedex) createPokemon(pokedata pokeapi.CatchResponse) Pokemon {
	return Pokemon{
		Name:       pokedata.Name,
		Experience: pokedata.BaseExperience,
		Stats: []struct {
			BaseStat int
			Effort   int
			Stat     struct {
				Name string
				URL  string
			}
		}(pokedata.Stats),
	}
}

func (p *Pokedex) AddPokemon(pokedata pokeapi.CatchResponse) {
	pokemon := p.createPokemon(pokedata)
	p.M.Lock()
	defer p.M.Unlock()
	p.Entries[pokemon.Name] = pokemon
}

func (p *Pokedex) ShowPokemonInfo(name string) (Pokemon, bool) {
	p.M.Lock()
	defer p.M.Unlock()
	pokemon, ok := p.Entries[name]
	return pokemon, ok
}
func (p *Pokedex) ListPokemon() {
	p.M.Lock()
	defer p.M.Unlock()
	for pokemon := range p.Entries {
		fmt.Println(pokemon)
	}
}

func (p *Pokedex) HasPokemon(name string) bool {
	p.M.Lock()
	defer p.M.Unlock()
	_, ok := p.Entries[name]
	return ok
}
