package grepkill

import (
	"log"

	"github.com/mitchellh/go-ps"
)

func GenerateProcess(done <-chan interface{}) <-chan ps.Process {
	process := make(chan ps.Process, 128)

	processList, err := ps.Processes()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer close(process)

		for _, p := range processList {
			select {
			case process <- p:
			case <-done:
				return
			}
		}
	}()

	return process
}
