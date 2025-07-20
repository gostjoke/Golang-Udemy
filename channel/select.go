package channel

import "fmt"

func SelectExample() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)
	receive(eve, odd, quit)

	fmt.Println("Exiting SelectExample")

	fmt.Println("Comma Start")
	comma()
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Received even:", v)
		case v := <-o:
			fmt.Println("Received odd:", v)
		case v := <-q:
			fmt.Println("Received quit:", v)
			return
		}
	}

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	close(e)
	close(o)
	q <- 0

}

func comma() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println("Value:", v, "Channel open:", ok)
	fmt.Println(<-c)
}
