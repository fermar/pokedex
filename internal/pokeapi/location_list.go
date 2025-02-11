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
	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// return nil, nil, err
	// }
	// if err != nil {
	// return nil, nil, err
	// }

	locList := LocAreaList{}
	err = json.Unmarshal(data, &locList)
	if err != nil {
		return LocAreaList{}, err
	}
	return locList, nil
	// for _, result := range locList.Results {
	// 	fmt.Println(result.Name)
	// }
	// fmt.Println()
	//
	// return locList.Next, locList.Previous, nil
}
