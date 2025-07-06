package main

import (
	"fmt"
	"scope/furtherexplored"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(furtherexplored.AValue)
	fmt.Println(furtherexplored.BValue)
	a, b := furtherexplored.ValueExchange()
	// a, b := furtherexplored.valueExchange() // this will not work because valueExchange is not exported

	fmt.Println(a, b)
}
