package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	mu := sync.Mutex{}
	cord := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		cord.L.Lock()
		for len(sharedRsc) == 0 {
			cord.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		cord.L.Unlock()

	}()

	cord.L.Lock()
	sharedRsc["rsc1"] = "foo"
	cord.Signal()
	cord.L.Unlock()

	wg.Wait()
}
