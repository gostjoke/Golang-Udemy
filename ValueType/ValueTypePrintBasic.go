package ValueType

import "fmt"

var y int

func ValueTypeF() {
	fmt.Println("y")
	z := 42
	z, b := 43, 11
	fmt.Printf("%v and %v", z, b)
	fmt.Printf("y = %d, z = %d, T = %T\n", y, z, b)
}
