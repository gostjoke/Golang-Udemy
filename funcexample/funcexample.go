package funcexample

import "fmt"

func ExampleFunction() {

	x := func(a int, b int) int {
		return a + b
	}

	result := x(2, 3)
	fmt.Println(result)

	y := Bar()
	fmt.Printf("This is an return function %T", y)
}

func Bar() func() int {
	// return a function
	return func() int {
		fmt.Println("This is an example function.")
		return 42
	}
}
