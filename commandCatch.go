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
	var prob float32
	chance = rand.Intn(pokeRes.BaseExperience)
	prob = float32(chance) / float32(pokeRes.BaseExperience)
	pokelog.Pl.Plogger.Printf("Experiencia: %v", pokeRes.BaseExperience)
	pokelog.Pl.Plogger.Printf("Chance: %v", chance)

	fmt.Printf("Throwing a Pokeball at %v...\n", param[0])
	pokelog.Pl.Plogger.Printf("la probabilidad es: %v\n", prob)
	if float32(prob) < 0.6 {
		fmt.Printf("%v se escapó\n", param[0])
	} else {
		fmt.Printf("%v atrapado!\n", param[0])
		UsrPokedex[param[0]] = pokeRes
		pokelog.Pl.Plogger.Printf("Se agrego %v a la pokedex.\n Cantidad de pokemones: %v\n", param[0], len(UsrPokedex))
	}

	// conf.next = locationsRes.Next
	// conf.previous = locationsRes.Previous

	// for _, pokenc := range pokeRes.PokemonEncounters {
	// 	fmt.Println(pokenc.Pokemon.Name)
	// }

	return nil
}
