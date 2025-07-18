package concurrency

import "fmt"

func doSomthing(x int) int {
	return x * 5
}

func DocumentationExample() {
	ch := make(chan int)
	go func() {
		ch <- doSomthing(5)
	}()
	fmt.Println(<-ch)
}
