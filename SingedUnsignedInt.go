package main

import "fmt"

func main() {
	var x uint8 = 255
	var y int8 = 127

	fmt.Printf("%v is of type %T, Bin: %b", x, x, x)
	fmt.Printf("\n%v is of type %T, Bin: %b", y, y, y)
}
