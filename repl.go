package main

import (
	"bufio"
	"fmt"
	"github.com/fermar/pokedex/internal/pokeapi"
	"log"
	"os"
	"strings"
)

type config struct {
	poqueapiClient pokeapi.Client
	next           *string
	previous       *string
}

func repl() {
	comandos := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	// var conf *config

	conf := config{}
	for {

		fmt.Print("Pockedex > ")

		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		input := cleanInput(scanner.Text())
		if len(input) < 1 {
			continue
		}
		if com, ok := comandos[input[0]]; !ok {
			fmt.Println("Unkown command")
		} else {
			err := com.callback(&conf)
			if err != nil {
				fmt.Println(err)
			}
		}
		// fmt.Printf("Your command was: %v \n", input[0])

	}

}

func cleanInput(text string) []string {

	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words

	// var ret []string
	// return ret

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// type locAreaList struct {
// 	Count    int     `json:"count"`
// 	Next     *string `json:"next"`
// 	Previous *string `json:"previous"`
// 	Results  []struct {
// 		Name string `json:"name"`
// 		URL  string `json:"url"`
// 	} `json:"results"`
// }

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Salir del pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Ayuda del pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "locations forward de Pokemon World",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "location backwards de Pokemon World",
			callback:    commandMapb,
		},
	}

}
