package RaceCondition

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func AtomicExample() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines before:", runtime.NumGoroutine())
	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) // 安全加1
			runtime.Gosched()            // 主動讓出時間片，讓其他 goroutine 有機會跑
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done() //這個 goroutine 結束，通知 WaitGroup。
		}()
		fmt.Println("Goroutine:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines after:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)

}
