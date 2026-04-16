package repl

import (
	"strings"
	"fmt"
	"os"
	"github.com/Khimich13/pokedex/internal/api"
)

type Config struct {
	Next *string
	Previous *string
}

type Command struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

var Commands map[string]Command
	
func init() {
	Commands = map[string]Command{
		"exit": {"exit", "Exit the Pokedex", commandExit},
		"help": {"help", "Displays a help message", commandHelp},
		"map":  {"map", "Show next 20 locations", commandMap},
		"mapb": {"mapb", "Show previous 20 locations", commandMapb},
	}
}

func CleanInput(text string) []string {
	output := strings.Fields(text)
	for i := range output {
		output[i] = strings.ToLower(output[i])
	}
	return output
}

func commandExit(state *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(state *Config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range Commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func commandMap(state *Config) error {
	if state.Next == nil && state.Previous != nil {
		fmt.Println("you're on the last page")
	} else {
		printGiven20(*state.Next, state)
	}
	return nil
}

func commandMapb(state *Config) error {
	if state.Previous == nil {
		fmt.Println("you're on the first page")
	} else {
		printGiven20(*state.Previous, state)
	}
	return nil
}

func printGiven20(url string, state *Config) {
	location, err := api.GetLocation(url)
	if err != nil{
		fmt.Println(err)
		return
	}
	for _, result := range location.Results {
		fmt.Println(result.Name)
	}
	state.Next = location.Next
	state.Previous = location.Previous
}