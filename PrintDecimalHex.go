// =======================================
// @File        : PrintDecimalHex.go
// @Author      : Tien-Wei Hsu
// @Date        : 2025-07-04 14:36
// @LastEditTime: 2025-07-04 14:36
// @Description : Description
// =======================================

package main

import "fmt"

func PrintDecimalHex() {
	adams := 42
	fmt.Printf("Binary: %b, Hexadecimal: %X\n", adams, adams)

	a, b, c, d, e, f := 1, 2, 3, 4, 5, 6
	fmt.Printf("Binary: %b, Hexadecimal: %X, HEX: %#X\n", a, a, a)
	fmt.Printf("Binary: %b, Hexadecimal: %X, HEX: %#X\n", b, b, b)
	fmt.Printf("Binary: %b, Hexadecimal: %X, HEX: %#X\n", c, c, c)
	fmt.Printf("Binary: %b, Hexadecimal: %X, HEX: %#X\n", d, d, d)
	fmt.Printf("Binary: %b, Hexadecimal: %X, HEX: %#X\n", e, e, e)
	fmt.Printf("Binary: %b, Hexadecimal: %X, HEX: %#X\n", f, f, f)

}

func main() {
	PrintDecimalHex()
}
