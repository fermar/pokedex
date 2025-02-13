package main

import (
	"fmt"
	"io"
	"os"

	"github.com/fermar/pokedex/internal/pokelog"
)

func commandTlog(conf *config) error {

	pokelog.Enabled = !pokelog.Enabled
	if pokelog.Enabled {
		pokelog.Logger.SetOutput(os.Stderr)
		fmt.Println("verbose activado")
	} else {
		pokelog.Logger.SetOutput(io.Discard)
		fmt.Println("verbose desactivado")
	}

	return nil
}
