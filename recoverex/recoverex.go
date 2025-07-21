package recoverex

import "fmt"

func RecoverExample() {
	f()
	fmt.Println("Return normally from f.")
}

func f() {
	defer func() {
		// recover is like python try except
		// it will catch the panic and prevent the program from crashing
		if r := recover(); r != nil {
			fmt.Println("Recovered in f:", r)
		}
	}()

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking in g.")
		panic(fmt.Sprintf("g(%d)", i))
	}
	defer fmt.Println("Defer in g:", i)
	fmt.Println("Calling g with", i+1)
	g(i + 1)
	fmt.Println("Returned normally from g with", i)
}
