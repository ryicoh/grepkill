package main

import (
	"fmt"
	"regexp"

	"github.com/mitchellh/go-ps"
	flag "github.com/spf13/pflag"
)

func main() {
	pattern := flag.StringP("pattern", "p", "", "pattern")
	flag.Parse()

	done := make(chan interface{})
	defer close(done)

	processCh := grepkill.generator(done)

	r := regexp.MustCompile(*pattern)
	matchedCh := grepkill.grep(done, processCh, r)

	matchedProcessList := []ps.Process{}
	for p := range matchedCh {
		matchedProcessList = append(matchedProcessList, p)
		fmt.Printf("%#v\n", p)
	}
}
