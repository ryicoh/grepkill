package grepkill

import (
	"regexp"

	"github.com/mitchellh/go-ps"
)

func Grep(done <-chan interface{}, process <-chan ps.Process, r *regexp.Regexp) <-chan ps.Process {
	matched := make(chan ps.Process, 128)

	go func() {
		defer close(matched)

		for p := range process {
			if !r.MatchString(p.Executable()) {
				continue
			}

			select {
			case matched <- p:
			case <-done:
				return
			}
		}
	}()

	return matched
}
