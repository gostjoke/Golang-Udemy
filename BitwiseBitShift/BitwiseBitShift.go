package main

import "fmt"

func BitwiseBitShift() {

	fmt.Printf("%d \t %b \n", 1, 1)
	fmt.Printf("%d \t %b \n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b \n", 1<<2, 1<<2)
}

func main() {
	BitwiseBitShift()
	fmt.Println("Bitwise Bit Shift Example Completed")
}
