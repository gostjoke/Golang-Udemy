package controlflow

import "fmt"

func ConditionExample() {
	// Example of if-else statement
	x := 10
	y := 20

	if x <= 10 && y >= 12 {
		fmt.Println("x is less than or equal to 10 and y is greater than or equal to 12")
	} else {
		fmt.Println("Condition not met")
	}

	if x >= 10 || y < 12 {
		fmt.Println("x is greater than 10 or y is less than 12")
	} else {
		fmt.Println("Neither condition met")
	}
}
