package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var comandos = map[string]cliCommand{}

//	var comandos = map[string]cliCommand{
//		"exit": {
//			name:        "exit",
//			description: "Salir del pokedex",
//			callback:    commandExit,
//		},
//		"help": {
//			name:        "help",
//			description: "Ayuda del pokedex",
//			callback:    commandHelp,
//		},
//	}
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, com := range comandos {
		fmt.Printf("%v: %v\n", com.name, com.description)
	}
	return nil
}

func Init() {
	var cliCom cliCommand
	cliCom.name = "exit"
	cliCom.description = "Salir del pokedex"
	cliCom.callback = commandExit
	comandos["exit"] = cliCom

	cliCom.name = "help"
	cliCom.description = "Ayuda del pokedex"
	cliCom.callback = commandHelp
	comandos["help"] = cliCom

}

func main() {
	// fmt.Println("Hello, World!")
	Init()
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pockedex > ")

		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		input := strings.Fields(strings.ToLower(scanner.Text()))
		if com, ok := comandos[input[0]]; !ok {
			fmt.Println("Unkown command")
		} else {
			com.callback()
		}
		// fmt.Printf("Your command was: %v \n", input[0])

	}
}
