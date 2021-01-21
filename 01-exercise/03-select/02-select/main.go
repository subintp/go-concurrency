package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "one"
	}()

	select {
	case m := <-ch:
		fmt.Println(m)
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out")
	}
}
