package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	// "io"
	// "net/http"
)

// func listLoc(url string) (next *string, prev *string, err error) {
//
//		res, err := http.Get(url)
//		if err != nil {
//			return nil, nil, err
//		}
//		body, err := io.ReadAll(res.Body)
//		res.Body.Close()
//		if res.StatusCode > 299 {
//			return nil, nil, err
//		}
//		if err != nil {
//			return nil, nil, err
//		}
//
//		locList := locAreaList{}
//		err = json.Unmarshal(body, &locList)
//		if err != nil {
//			return nil, nil, err
//		}
//
//		for _, result := range locList.Results {
//			fmt.Println(result.Name)
//		}
//		fmt.Println()
//
//		return locList.Next, locList.Previous, nil
//	}

func commandMap(conf *config) error {
	fmt.Println("Mostrando locations forward")
	locationsRes, err := conf.poqueapiClient.ListLoc(conf.next)
	if err != nil {
		return err
	}
	conf.next = locationsRes.Next
	conf.previous = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(conf *config) error {
	fmt.Println("Mostrando locations backwards")
	if conf.previous == nil {
		return errors.New("you're ont the first page")
	}
	locationsRes, err := conf.poqueapiClient.ListLoc(conf.previous)
	if err != nil {
		return err
	}
	conf.next = locationsRes.Next
	conf.previous = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// func commandMapb(conf *config) error {
// 	fmt.Println("Mostrando locations forward")
// 	var url string
// 	if conf.previous == nil {
// 		fmt.Println("you're on the fist page")
// 		url = "https://pokeapi.co/api/v2/location-area/"
// 	} else {
// 		url = *conf.previous
// 	}
// 	var err error
// 	conf.next, conf.previous, err = listLoc(url)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
