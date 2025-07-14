package funcexample

import (
	"fmt"
)

func Fibonacci(a int, b int, Stop int) int {
	fmt.Printf("%v \n", a+b)
	if a+b > Stop {
		return a + b
	}
	return Fibonacci(b, a+b, Stop)
}
