package RaceCondition

import (
	"fmt"
	"sync"
	"time"
)

func EditValue(s string, value *int, wg *sync.WaitGroup) {
	defer wg.Done()
	if s == "+" {
		(*value)++
	} else if s == "-" {
		(*value)--
	} else {
		fmt.Println("Invalid operation")
		return
	}
	time.Sleep(10 * time.Millisecond) // Simulate some processing time
	fmt.Println("Value after operation:", *value)

	fmt.Println("Value after increment:", *value)
}

func main() {
	var wg sync.WaitGroup
	Value := 10000
	for i := 0; i < 10000; i++ {
		wg.Add(4)
		go EditValue("+", &Value, &wg)
		go EditValue("-", &Value, &wg)
		go EditValue("+", &Value, &wg)
		go EditValue("-", &Value, &wg)
	}
	fmt.Println("Value before wait:", Value)
	wg.Wait()
}
