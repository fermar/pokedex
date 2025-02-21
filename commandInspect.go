package main

import (
	"errors"
	"fmt"
	// "math/rand"

	"github.com/fermar/pokedex/internal/pokeapi"
	// "github.com/fermar/pokedex/internal/pokelog"
)

// "fmt"
// "github.com/fermar/pokedex/internal/pokelog"

func commandInspect(conf *config, param []string) error {
	if len(param) < 1 {
		return errors.New("falta parametro, nombre del pokemon")
	}

	var pokeData pokeapi.PokemonInfo
	var ok bool
	pokeData, ok = UsrPokedex[param[0]]
	if !ok {
		fmt.Printf("No has atrapado a %v aún\n", param[0])
		return nil
	}
	fmt.Printf("Nombre: %v\n", pokeData.Name)
	fmt.Printf("Heigth: %v\n", pokeData.Height)
	fmt.Printf("Weight: %v\n", pokeData.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokeData.Stats {
		fmt.Printf("\t-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types: ")
	for _, tipo := range pokeData.Types {
		fmt.Printf("\t- %v\n", tipo.Type.Name)
	}

	// pokeRes, err := conf.poqueapiClient.Catch(param[0])
	// if err != nil {
	// 	return err
	// }
	// var chance int
	// chance = rand.Intn(pokeRes.BaseExperience)
	// prob := chance / pokeRes.BaseExperience
	// pokelog.Pl.Plogger.Printf("Experiencia: %v", pokeRes.BaseExperience)
	// pokelog.Pl.Plogger.Printf("Chance: %v", chance)
	//
	// fmt.Printf("Throwing a Pokeball at %v...\n", param[0])
	//
	// if float32(prob) > 0.6 {
	// 	fmt.Printf("%v se escapó\n", param[0])
	// } else {
	// 	fmt.Printf("%v atrapado!\n", param[0])
	// 	UsrPokedex[param[0]] = pokeRes
	// 	pokelog.Pl.Plogger.Printf("Se agrego %v a la pokedex.\n Cantidad de pokemones: %v\n", param[0], len(UsrPokedex))
	// }
	//
	// // conf.next = locationsRes.Next
	// // conf.previous = locationsRes.Previous
	//
	// // for _, pokenc := range pokeRes.PokemonEncounters {
	// // 	fmt.Println(pokenc.Pokemon.Name)
	// // }
	//
	return nil
}
