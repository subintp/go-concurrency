package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	deadline := time.Now().Add(500 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			// Simulate work.
			time.Sleep(50 * time.Millisecond)
			// Report result.
			ch <- data{"123"}
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ch := compute(ctx)

	select {
	case d := <-ch:
		fmt.Printf("work complete: %s\n", d)
	case <-ctx.Done():
		fmt.Printf("work incomplete")
	}
}
