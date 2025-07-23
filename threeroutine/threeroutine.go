package threeroutine

import (
	"fmt"
	"sync"
)

// if we have 3 routines, need to do order t1 t2 t3
func ThreeroutineEX() {
	var wg sync.WaitGroup
	wg.Add(3)

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		defer wg.Done()
		// Wait for T1 to signal before running T2
		<-ch1
		fmt.Println("T2 running")
		ch2 <- struct{}{}
	}()

	go func() {
		defer wg.Done()
		// Wait for T2 to signal before running T3
		<-ch2
		fmt.Println("T3 running")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("T1 running")
		ch1 <- struct{}{}
	}()

	wg.Wait()
}
