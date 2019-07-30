package main

import (
  "fmt"
  "github.com/mitchellh/go-ps"
  "math"
)

func main() {

  find := func(done <-chan interface{}, id <-chan interface{}) <-chan ps.Process {
    p := make(chan ps.Process)
    go func() {
      for i range id {
        p, err := ps.FindProcess(i)
        if err != nil {
          log.Fatal(err)
        }

        if p != nil {
          fmt.Println(p, err)
        }
      }
    }
  }

  done := make(chan interface{})
  defer close(done)

  id := make(chan int)
  p := find(done, id)

  for i := 0; i < int(math.Pow(2, 16)); i++ {
    id := 
  }
}
