package commands

import (
	"fmt"

	"github.com/2bitburrito/pokedex-cli/internal/types"
)

func CommandHelp(*types.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
