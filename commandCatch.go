package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/fermar/pokedex/internal/pokelog"
)

// "fmt"
// "github.com/fermar/pokedex/internal/pokelog"

func commandCatch(conf *config, param []string) error {
	if len(param) < 1 {
		return errors.New("falta parametro, nombre del pokemon")
	}

	pokeRes, err := conf.poqueapiClient.Catch(param[0])
	if err != nil {
		return err
	}
	var chance int
	chance = rand.Intn(101)
	pokelog.Pl.Plogger.Printf("Experiencia: %v", pokeRes.BaseExperience)
	pokelog.Pl.Plogger.Printf("Chance: %v", chance)

	fmt.Printf("Throwing a Pokeball at %v", param[0])

	if pokeRes.BaseExperience > chance {
		fmt.Printf("%v se escapó", param[0])
	} else {
		fmt.Printf("%v atrapado!", param[0])
	}

	// conf.next = locationsRes.Next
	// conf.previous = locationsRes.Previous

	// for _, pokenc := range pokeRes.PokemonEncounters {
	// 	fmt.Println(pokenc.Pokemon.Name)
	// }

	return nil
}
