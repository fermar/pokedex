package main

import (
	"time"

	"github.com/fermar/pokedex/internal/pokeapi"
	"github.com/fermar/pokedex/internal/pokelog"
)

func main() {
	// fmt.Println("Hello, World!")
	pokelog.StartPokelogger(false)
	pokeClient := pokeapi.NewClient(5*time.Second, 20*time.Second)
	conf := &config{
		poqueapiClient: pokeClient,
	}
	repl(conf)

}
