package main

import (
	"context"
	"fmt"
)

func main() {
	generator := func(ctx context.Context) <-chan int {
		out := make(chan int)

		go func() {
			defer close(out)
			for i := 1; i < 10; i++ {
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			}
		}()

		return out
	}

	ctx, cancel := context.WithCancel(context.Background())

	ch := generator(ctx)

	for n := range ch {
		fmt.Println(n)
		if n == 5 {
			cancel()
		}
	}
}
