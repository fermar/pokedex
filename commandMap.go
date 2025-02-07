package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func listLoc(url string) (next string, prev string, err error) {

	res, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return "", "", err
	}
	if err != nil {
		return "", "", err
	}

	locList := locAreaList{}
	err = json.Unmarshal(body, &locList)
	if err != nil {
		return "", "", err
	}

	for _, result := range locList.Results {
		fmt.Println(result.Name)
	}
	fmt.Println()

	return string(locList.Next), string(locList.Previous), nil
}

func commandMap(conf *config) error {
	fmt.Println("Mostrando locations forward")
	var url string
	if conf.next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = conf.next
	}
	var err error
	conf.next, conf.previous, err = listLoc(url)
	if err != nil {
		return err
	}

	return nil
}

func commandMapb(conf *config) error {
	fmt.Println("Mostrando locations forward")
	var url string
	if conf.previous == "" {
		fmt.Println("you're on the fist page")
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = conf.previous
	}
	var err error
	conf.next, conf.previous, err = listLoc(url)
	if err != nil {
		return err
	}

	return nil
}
