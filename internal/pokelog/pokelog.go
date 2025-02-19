package pokelog

import (
	"io"
	"log"
	"os"
)

// type PokeLogger struct {
// 	Plogger *log.Logger
// 	Enabled bool
// }

var Plogger *log.Logger
var Enabled bool

// var Pl PokeLogger{}

func StartPokelogger(startEnabled bool) {

	// Pl := PokeLogger{Enabled: startEnabled,
	// 	Plogger: log.New(io.Discard, "Pokelog", log.LstdFlags),
	// }
	// Pl.enabled = startEnabled
	// Pl.Plogger = log.New(io.Discard, "PokeLog:", log.LstdFlags)
	Plogger = log.New(io.Discard, "PokeLog:", log.LstdFlags)
	Enabled = startEnabled
	if Enabled {
		EnLog()
	} else {
		DisLog()
	}
	// if Pl.Enabled {
	// 	Pl.EnLog()
	// } else {
	// 	Pl.DisLog()
	// }
}
func EnLog() {
	Enabled = true
	Plogger.SetOutput(os.Stderr)
}

//	func (pl *PokeLogger) EnLog() {
//		pl.Enabled = true
//		pl.Plogger.SetOutput(os.Stderr)
//	}
func DisLog() {
	Enabled = false
	Plogger.SetOutput(io.Discard)
} // func (pl *PokeLogger) DisLog() {
// 	pl.Enabled = false
// 	pl.Plogger.SetOutput(io.Discard)
// }
