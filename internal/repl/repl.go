package repl

import (
	"strings"
	"fmt"
	"os"
	"github.com/Khimich13/pokedex/internal/api"
	"math/rand"
)

type Config struct {
	Next *string
	Previous *string
}

type Command struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		"exit":    {"exit",    "Exit the Pokedex",        commandExit},
		"help":    {"help",    "Displays a help message", commandHelp},
		"map":     {"map",     "Show next 20 locations",  commandMap},
		"mapb":    {"mapb",    "Show prev 20 locations",  commandMapb},
		"explore": {"explore", "Show location info",      commandExplore},
		"catch":   {"catch",   "Try to catch a Pokemon",  commandCatch},
		"inspect": {"inspect", "Show Pokemon stats",      commandInspect},
		"pokedex": {"pokedex", "Show caught Pokemons",    commandPokedex},
	}
}

var pokedex = map[string]api.Pokemon{}

const maxExp int = 650

func CleanInput(text string) []string {
	output := strings.Fields(text)
	for i := range output {
		output[i] = strings.ToLower(output[i])
	}
	return output
}

func commandExit(state *Config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(state *Config, params []string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range Commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func commandMap(state *Config, params []string) error {
	if state.Next == nil && state.Previous != nil {
		fmt.Println("you're on the last page")
	} else if state.Next == nil {
		printGiven20(api.LocationAreaUrl, state)
	} else {
		printGiven20(*state.Next, state)
	}
	return nil
}

func commandMapb(state *Config, params []string) error {
	if state.Previous == nil {
		fmt.Println("you're on the first page")
	} else {
		printGiven20(*state.Previous, state)
	}
	return nil
}

func commandExplore(state *Config, params []string) error {
	if len(params) == 0 {
		fmt.Println("you need to provide a valid location id or name")
	} else {
		printExploringResults(params[0])
	}
	return nil
}

func commandCatch(state *Config, params []string) error {
	if len(params) == 0 {
		fmt.Println("you need to provide a valid Pokemon name")
	} else {
		printTryCatch(params[0])
	}
	return nil
}

func commandInspect(state *Config, params []string) error {
	if len(params) == 0 {
		fmt.Println("you need to provide a valid Pokemon name")
	} else {
		printPokemonStats(params[0])
	}
	return nil
}

func commandPokedex(state *Config, params []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}
	return nil
}

func printExploringResults(area string) {
	location, err := api.GetData[api.LocationArea](api.LocationAreaUrl + area)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Exploring " + area + "...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
}

func printGiven20(url string, state *Config) {
	location, err := api.GetData[api.LocationArea](url)
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

func printTryCatch(pokemonName string) {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := api.GetData[api.Pokemon](api.PokemonUrl + pokemonName)
	if err != nil{
		fmt.Println(err)
		return
	}
	chance := rand.Intn(maxExp)
	if pokemon.BaseExperience <= chance {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		pokedex[pokemonName] = *pokemon
	} else {
		fmt.Printf("%s was escaped!\n", pokemonName)
	}
}

func printPokemonStats(pokemonName string) {
	if pokemon, exist := pokedex[pokemonName]; !exist {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %v\n",   pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("  - %v\n", pokemonType.Type.Name)
		}
	}
}