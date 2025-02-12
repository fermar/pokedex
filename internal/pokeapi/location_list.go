package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/fermar/pokedex/internal/pokecache"
)

func (cli *Client) ListLoc(pageURL *string, cache *pokecache.Cache, logger *log.Logger) (LocAreaList, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	logger.Printf("Buscando URL: %v", url)

	data, hit := cache.Get(url)

	if !hit {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocAreaList{}, err
		}

		res, err := cli.httpClient.Do(req)
		if err != nil {
			return LocAreaList{}, err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocAreaList{}, err
		}
		cache.Add(url, data)
	}

	locList := LocAreaList{}
	err := json.Unmarshal(data, &locList)
	if err != nil {
		return LocAreaList{}, err
	}
	return locList, nil

}
