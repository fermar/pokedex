package main

import (
	"errors"
	"fmt"

	"github.com/fermar/pokedex/internal/pokelog"
)

func commandMap(conf *config) error {
	pokelog.Logger.Println("Mostrando locations forward")
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
	pokelog.Logger.Println("Mostrando locations backwards")
	if conf.previous == nil {
		return errors.New("you're on the first page")
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
