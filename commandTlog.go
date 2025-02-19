package main

import (
	"fmt"
	// "io"
	// "os"

	"github.com/fermar/pokedex/internal/pokelog"
)

func commandTlog(conf *config) error {

	pokelog.Enabled = !pokelog.Enabled
	if pokelog.Enabled {
		// pokelog.Plogger.SetOutput(os.Stderr)
		pokelog.EnLog()
		fmt.Println("verbose activado")
	} else {
		// pokelog.Plogger.SetOutput(io.Discard)
		pokelog.DisLog()
		fmt.Println("verbose desactivado")
	}

	return nil
}
