package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (cli *Client) ListLoc(pageURL *string) (LocAreaList, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// res, err := http.Get(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocAreaList{}, err
	}

	res, err := cli.httpClient.Do(req)
	if err != nil {
		return LocAreaList{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocAreaList{}, err
	}
	locList := LocAreaList{}
	err = json.Unmarshal(data, &locList)
	if err != nil {
		return LocAreaList{}, err
	}
	return locList, nil

}
