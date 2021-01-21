package main

import "fmt"

func main() {

	consumer := func(ch <-chan int) {
		for v := range ch {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("Done receiving!")
	}

	owner := func() <-chan int {
		ch := make(chan int, 6)

		go func() {
			defer close(ch)

			for i := 0; i < 6; i++ {
				ch <- i
			}
		}()

		return ch
	}

	ch := owner()
	consumer(ch)
}
