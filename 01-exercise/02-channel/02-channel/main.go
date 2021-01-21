package main

import "fmt"

func main() {
	channel := make(chan int, 6)
	go func() {
		for i := 0; i < 6; i++ {
			channel <- i
		}
		close(channel)
	}()

	for n := range channel {
		fmt.Println(n)
	}
}
