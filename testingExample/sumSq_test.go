package testingExample

import "testing"

func TestSumsq(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	//  table-driven test（表格驅動測試）
	// 這是一種常見的 Go 語言測試模式，使用結構體來定義測試用例，
	// 使得測試更易於擴展和維護。
	tests := []test{
		{[]int{2, 3}, 25},
		{[]int{3, 4}, 49},
		{[]int{5, 6}, 111},
		{[]int{7, 8}, 125},
	}

	// x := Sumsq(3, 4)
	// if x != 49 {
	// t.Errorf("Sumsq(3, 4) = %d; want 49", x)
	// }

	for _, v := range tests {
		x := Sumsq(v.data[0], v.data[1])
		if x != v.answer {
			t.Errorf("Sumsq(%d, %d) = %d; want %d", v.data[0], v.data[1], x, v.answer)
		}
	}
	// go test -v
	// go test -cover
	// coverage: 100.0% of statements
	// go tool cover -html=coverage.out
}
