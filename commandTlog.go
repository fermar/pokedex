package main

import (
	"fmt"
	"io"
	"os"

	"github.com/fermar/pokedex/internal/pokelog"
)

func commandTlog(conf *config) error {

	pokelog.Pl.Enabled = !pokelog.Pl.Enabled
	if pokelog.Pl.Enabled {
		pokelog.Pl.Plogger.SetOutput(os.Stderr)
		fmt.Println("verbose activado")
	} else {
		pokelog.Pl.Plogger.SetOutput(io.Discard)
		fmt.Println("verbose desactivado")
	}

	return nil
}
