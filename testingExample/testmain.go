package testingExample

func Fibonacci(a int, b int, Stop int) int {
	if a+b > Stop {
		return a + b
	}
	return Fibonacci(b, a+b, Stop)
}
