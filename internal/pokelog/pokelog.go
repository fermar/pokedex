package pokelog

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger
var Enabled bool

func StartPokelogger() {
	Enabled = false
	Logger = log.New(io.Discard, "PokeLog:", log.LstdFlags)
}

func EnLog() {
	Enabled = true
	Logger.SetOutput(os.Stderr)
}

func DisLog() {
	Enabled = false
	Logger.SetOutput(io.Discard)
}

