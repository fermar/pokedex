package main

import (
	"fmt"
)

func commandHelp() error {
	comandos := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, com := range comandos {
		fmt.Printf("%v: %v\n", com.name, com.description)
	}
	return nil
}
