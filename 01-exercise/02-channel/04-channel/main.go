package main

import "fmt"

func genMsg(ch1 chan<- string) {
	ch1 <- "Hello World"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	ch2 <- <-ch1
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	defer close(ch1)
	defer close(ch2)

	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	relayedMessage := <-ch2
	fmt.Println(relayedMessage)
}
