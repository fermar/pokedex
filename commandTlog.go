package main

import (
	"fmt"
	"io"
	"os"
)

func commandTlog(conf *config) error {
	conf.enLog = !conf.enLog
	if conf.enLog {
		conf.logger.SetOutput(os.Stderr)
		fmt.Println("verbose activado")
	} else {
		conf.logger.SetOutput(io.Discard)
		fmt.Println("verbose DESactivado")
	}

	return nil
}
