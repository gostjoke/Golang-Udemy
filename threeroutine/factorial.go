package threeroutine

import "fmt"

// 計算 n 的階乘，並回傳 int
func Factorial(n int) int {
	out := make(chan int)
	go func() {
		total := 1
		for i := 1; i <= n; i++ {
			total *= i
		}
		out <- total
		close(out)
	}()
	return <-out // 等待 goroutine 結果
}

// 包裝一個對外使用的函數
func Factorialex(n int) {
	result := Factorial(n)
	fmt.Printf("Factorial of %d is %d\n", n, result)
}
