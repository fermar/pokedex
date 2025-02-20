package main

import (
	"errors"
	"fmt"
)

// "fmt"
// "github.com/fermar/pokedex/internal/pokelog"

func commandExplore(conf *config, param []string) error {
	if len(param) < 1 {
		return errors.New("falta parametro a explorar")
	}

	pokeRes, err := conf.poqueapiClient.Explore(param[0])
	if err != nil {
		return err
	}
	// conf.next = locationsRes.Next
	// conf.previous = locationsRes.Previous

	for _, pokenc := range pokeRes.PokemonEncounters {
		fmt.Println(pokenc.Pokemon.Name)
	}

	return nil
}
