package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	// "github.com/fermar/pokedex/internal/pokecache"
	"github.com/fermar/pokedex/internal/pokelog"
)

func (cli *Client) Explore(area string) (PokemonesList, error) {

	url := baseURL + "/location-area/" + area
	// if pageURL != nil {
	// 	url = *pageURL
	// }

	pokelog.Pl.Plogger.Printf("Buscando URL: %v", url)

	data, hit := cli.cache.Get(url)

	if !hit {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonesList{}, err
		}

		res, err := cli.httpClient.Do(req)
		if err != nil {
			return PokemonesList{}, err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonesList{}, err
		}
		cli.cache.Add(url, data)
	}

	pokeList := PokemonesList{}
	err := json.Unmarshal(data, &pokeList)
	if err != nil {
		return PokemonesList{}, err
	}
	return pokeList, nil

}
