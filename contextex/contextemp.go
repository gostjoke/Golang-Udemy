package contextex

import (
	"context"
	"fmt"
)

func ContextExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure the context is cancelled to avoid leaks
	fmt.Println("Context Example started")
	for v := range gen(ctx) {
		println(v)
		if v == 5 {
			break
		}
	}
	fmt.Println("Context Example finished")

}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
