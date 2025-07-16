package pointer

import "fmt"

func Dereferencing() {
	a := 3.0
	fmt.Println("Original value:", a)
	fmt.Println("Address of a:", &a)
	// fmt.Println("Dereferenced value:", *(&a))
	fmt.Printf("type of %T", &a)
	// *(&a) = 5.0
	fmt.Println("\nModified value:", a)

}
