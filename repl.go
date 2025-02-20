package main

import (
	"bufio"
	"fmt"

	"github.com/fermar/pokedex/internal/pokeapi"
	"github.com/fermar/pokedex/internal/pokelog"
	"log"
	"os"
	"strings"
)

type config struct {
	poqueapiClient pokeapi.Client
	next           *string
	previous       *string
}

func repl(conf *config) {
	comandos := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {

		if pokelog.Pl.Enabled {
			fmt.Print("Pockedex (verbose)> ")
		} else {
			fmt.Print("Pockedex > ")
		}
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

			err := com.callback(conf, input[1:])
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}

func cleanInput(text string) []string {

	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

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
		"explore": {
			name:        "explore",
			description: "Explorar área del pokeworld",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Atrapar un pokemon",
			callback:    commandCatch,
		},
		"tlog": {
			name:        "tlog",
			description: "Toggle verbose output",
			callback:    commandTlog,
		},
	}

}
