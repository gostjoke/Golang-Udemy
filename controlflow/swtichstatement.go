package controlflow

import (
	"fmt"
)

func SwitchStatementExample() {
	// Example of switch statement
	x := 10

	switch x {
	case 10:
		println("x is 10")
	case 20:
		println("x is 20")
	default:
		println("x is neither 10 nor 20")
	}

	// Example of switch with multiple cases
	switch x {
	case 5, 10, 15:
		println("x is either 5, 10, or 15")
	default:
		println("x is something else")
	}

	switch x := 1; x {
	case 1:
		fmt.Println("First")
		fallthrough
	case 2:
		fmt.Println("Second")
	}

}
