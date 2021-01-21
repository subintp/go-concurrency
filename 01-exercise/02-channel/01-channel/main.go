package main

import "fmt"

func main() {
	channel := make(chan int)

	go func(a, b int) {
		c := a + b
		channel <- c
	}(1, 2)

	sum := <-channel
	fmt.Printf("computed value %v\n", sum)
}
