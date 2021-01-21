package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)

	wg.Add(1)
	go func() {
		defer wg.Done()

		cond.L.Lock()
		for len(sharedRsc) == 0 {
			cond.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		cond.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		cond.L.Lock()
		for len(sharedRsc) == 0 {
			cond.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		cond.L.Unlock()
	}()

	// writes changes to sharedRsc
	cond.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	cond.Broadcast()
	cond.L.Unlock()

	wg.Wait()
}
