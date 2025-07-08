package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func ChannelStateExample() {
	// Example of channel state
	rand.Seed(time.Now().UnixNano())
	d1 := time.Duration(rand.Intn(10))
	d2 := time.Duration(rand.Intn(10))
	ch1 := make(chan int)
	ch2 := make(chan int)
	// Sending a value to the channel
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 42
	}()
	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 41
	}()

	select {
	case val1 := <-ch1:
		fmt.Println("Received from ch1:", val1)
	case val2 := <-ch2:
		fmt.Println("Received from ch2:", val2)
	}

	// Closing the channel
	close(ch1)
	close(ch2)

}
