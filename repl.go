package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func repl() {
	comandos := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

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
			com.callback()
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
	callback    func() error
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
	}

}
