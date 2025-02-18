package pokelog

import (
	"io"
	"log"
	"os"
)

type PokeLogger struct {
	Plogger *log.Logger
	Enabled bool
}

var Pl PokeLogger

func StartPokelogger(startEnabled bool) {

	Pl := PokeLogger{Enabled: startEnabled}
	// Pl.enabled = startEnabled
	Pl.Plogger = log.New(io.Discard, "PokeLog:", log.LstdFlags)
	if Pl.Enabled {
		Pl.EnLog()
	} else {
		Pl.DisLog()
	}
}

func (pl *PokeLogger) EnLog() {
	pl.Enabled = true
	pl.Plogger.SetOutput(os.Stderr)
}

func (pl *PokeLogger) DisLog() {
	pl.Enabled = false
	pl.Plogger.SetOutput(io.Discard)
}
