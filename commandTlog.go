package main

import (
	"fmt"

	"github.com/fermar/pokedex/internal/pokelog"
	// "io"
	// "os"
	// "github.com/fermar/pokedex/internal/pokelog"
)

func commandTlog(conf *config) error {

	// pokelog.Enabled = !pokelog.Enabled
	pokelog.Pl.Enabled = !pokelog.Pl.Enabled
	// if pokelog.Enabled {
	if pokelog.Pl.Enabled {
		// pokelog.Plogger.SetOutput(os.Stderr)
		// pokelog.EnLog()
		pokelog.Pl.EnLog()
		fmt.Println("verbose activado")
	} else {
		// pokelog.Plogger.SetOutput(io.Discard)
		pokelog.Pl.DisLog()
		fmt.Println("verbose desactivado")
	}

	return nil
}
