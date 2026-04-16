package main

import (
	"strings"
	"fmt"
	"os"
)

type config struct {
	Next *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commands map[string]cliCommand
	
func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: 	 commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of 20 Pokemon world locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of 20 previous Pokemon world locations",
			callback: commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.Fields(text)
	for i := range output {
		output[i] = strings.ToLower(output[i])
	}
	return output
}

func commandExit(state *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(state *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(state *config) error {
	if state.Next == nil && state.Previous != nil {
		fmt.Println("you're on the last page")
	} else {
		printGiven20(*state.Next, state)
	}
	return nil
}

func commandMapb(state *config) error {
	if state.Previous == nil {
		fmt.Println("you're on the first page")
	} else {
		printGiven20(*state.Previous, state)
	}
	return nil
}