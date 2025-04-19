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
	Height     int
	Weight     int

	Stats []struct {
		BaseStat int
		Effort   int
		Stat     struct {
			Name string
			URL  string
		}
	}
	Types []struct {
		Slot int
		Type struct {
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
		Height:     pokedata.Height,
		Weight:     pokedata.Weight,

		Stats: []struct {
			BaseStat int
			Effort   int
			Stat     struct {
				Name string
				URL  string
			}
		}(pokedata.Stats),
		Types: []struct {
			Slot int
			Type struct {
				Name string
				URL  string
			}
		}(pokedata.Types),
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
	dexLength := len(p.Entries)
	if dexLength < 1 {
		fmt.Println("0 Pokemon found")
		return
	}
	fmt.Printf("You have caught: %d Pokemon\n", dexLength)
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
func (p *Pokedex) Inspect(name string) {
	p.M.Lock()
	defer p.M.Unlock()
	pokemon, ok := p.Entries[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, el := range pokemon.Types {
		fmt.Printf("-%v\n", el.Type.Name)
	}
}
