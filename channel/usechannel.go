package channel

import "fmt"

func UseChannelExample() {
	// send
	c := make(chan int, 2)

	go foo(c)
	// receive
	bar(c)
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

// send
func foo(c chan<- int) {

	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println("Received value from channel:", <-c)
}

// will cause deadlock if not closed
func RangeChannelExample() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// close(ch)
	}()

	for val := range ch {
		fmt.Println("Received value:", val)
	}
	fmt.Println("Channel closed, exiting range loop")

}
