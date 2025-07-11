package slicefun

import "fmt"

func Makefun() {
	a := make([]int, 5)
	a = append(a, 1, 2, 3, 4, 5)
	fmt.Println(a)
	b := []int{1, 2, 3, 4, 5} // 初始化陣列
	c := make([]int, 5)       // 預留空間為 5
	copy(c, b)                // 將 c 的內容複製到 b
	fmt.Println("b value:", c)

	fmt.Println("b cap:", cap(b))
	b = append(b, 6, 7, 8, 9, 10) // 添加元素到 b
	fmt.Println("b cap:", cap(b))
	d := make([]int, 0, 5)                       // 預留空間為 5
	d = append(d, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10) // 添加元素到 d
	fmt.Println("d value:", d)
}
