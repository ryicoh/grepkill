package main

import (
	"fmt"
	"regexp"

	"github.com/mitchellh/go-ps"
	"github.com/ryicoh/grepkill"
	flag "github.com/spf13/pflag"
)

func main() {
	pattern := flag.StringP("pattern", "p", "", "pattern")
	flag.Parse()

	done := make(chan interface{})
	defer close(done)

	processCh := grepkill.GenerateProcess(done)

	r := regexp.MustCompile(*pattern)
	matchedCh := grepkill.Grep(done, processCh, r)

	matchedProcessList := []ps.Process{}
	for p := range matchedCh {
		matchedProcessList = append(matchedProcessList, p)
		fmt.Printf("%#v\n", p)
	}

	if len(matchedProcessList) <= 0 {
		fmt.Printf("no match '%s'\n", *pattern)
		return
	}

	if !grepkill.AskForConfirmation() {
		fmt.Println("cancelled by user")
		return
	}

	for _, p := range matchedProcessList {
		grepkill.Kill(p.Pid())
		fmt.Printf("killed PID %d, Name %s\n", p.Pid(), p.Executable())
	}
}
