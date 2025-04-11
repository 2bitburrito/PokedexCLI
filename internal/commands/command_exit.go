package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
)

func CommandExit(*pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Failed to exit gracefully - Please close with ctrl-c")
}
