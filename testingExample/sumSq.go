package testingExample

import "fmt"

func Sumsq(a, b int) int {
	fmt.Println("Sumsq of ", a, "and", b)
	return a*a + 2*a*b + b*b
}
