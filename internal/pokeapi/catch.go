package pokeapi

import (
	"encoding/json"
	"errors"
	// "fmt"
	"io"
	"net/http"

	// "github.com/fermar/pokedex/internal/pokecache"
	"github.com/fermar/pokedex/internal/pokelog"
)

func (cli *Client) Catch(pokeName string) (PokemonInfo, error) {

	url := baseURL + "/pokemon/" + pokeName
	// if pageURL != nil {
	// 	url = *pageURL
	// }

	pokelog.Pl.Plogger.Printf("Buscando URL: %v", url)

	data, hit := cli.cache.Get(url)

	if !hit {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonInfo{}, err
		}

		res, err := cli.httpClient.Do(req)
		if err != nil {
			return PokemonInfo{}, err
		}
		defer res.Body.Close()
		if res.StatusCode == 404 {
			// fmt.Printf("El pokemon %v no existe\n", pokeName)
			return PokemonInfo{}, errors.New("pokemon Inexistente")
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonInfo{}, err
		}
		cli.cache.Add(url, data)
	}

	pokeInfo := PokemonInfo{}
	err := json.Unmarshal(data, &pokeInfo)
	if err != nil {
		return PokemonInfo{}, err
	}
	return pokeInfo, nil

}
