package Iota

import "fmt"

const (
	d0 = 1 << iota // 1 << 0 = 1
	d1
	d2
	d3
)

func Iota() {
	const (
		c0 = iota // 0
		c1        // 1
		c2        // 2
		c3        // 3
	)

	fmt.Println(c0, c1, c2, c3)
	fmt.Println(d0, d1, d2, d3)
}
