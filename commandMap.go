package main

import (
	"errors"
	"fmt"
)

func commandMap(conf *config) error {
	fmt.Println("Mostrando locations forward")
	locationsRes, err := conf.poqueapiClient.ListLoc(conf.next, conf.cache, conf.logger)
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
	locationsRes, err := conf.poqueapiClient.ListLoc(conf.previous, conf.cache, conf.logger)
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
