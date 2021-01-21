package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}
	}()

	for i := 0; i < 3; {
		select {
		case m := <-ch:
			i++
			fmt.Println(m)
		default:
			fmt.Println("processing...")
			time.Sleep(1 * time.Second)
		}
	}
}
