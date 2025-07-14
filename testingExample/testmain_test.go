package testingExample

import (
	"log"
	"testing"
)

// go test -v or go test
// go test -run TestPuppyBark
func TestPuppyBark(t *testing.T) {
	got := Fibonacci(0, 1, 10)
	want := 13

	if got != want {
		t.Errorf("Fibonacci(0, 1, 10) = %d; want %d", got, want)
		log.Fatalf("Fibonacci(0, 1, 10) = %d; want %d", got, want)
	}
}

// go test -bench=.
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(0, 1, 1000) // 執行你要測效能的程式
	}
}
