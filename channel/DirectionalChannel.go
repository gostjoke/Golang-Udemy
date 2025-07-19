package channel

import "fmt"

func DirectionalChannelExample() {
	ch := make(chan int, 2)
	// fmt.Println("Sent values to directional channel:", <-ch) will cause a compile-time error

	ch <- 42
	ch <- 43
	fmt.Println("Sent values to directional channel:", <-ch) // This will cause a compile-time error
	fmt.Println("Sent values to directional channel:", <-ch)

	fmt.Println("------")
	fmt.Println("Directional channel example completed")
}
