package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func FanoutExample() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		println("Received value:", v)
	}

	fmt.Println("Fanout example completed")
}

// populate 向 c1 發送 0 到 99 的整數
func populate(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// FanoutIn 接收來自 c1 的值，並將處理結果發送到 c2
// 使用 goroutine 處理每個值，並在所有 goroutine 完成後
func fanOutIn(c1 <-chan int, c2 chan<- int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			defer wg.Done()
			// 模擬工作
			result := timeConsumingWork(v2)
			c2 <- result
		}(v)
	}

	// 在 goroutine 裡等所有 worker 做完，再關閉 c2
	go func() {
		wg.Wait()
		close(c2)
	}()
}

// 模擬耗時工作
// 這裡可以替換成實際的工作邏輯
func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
