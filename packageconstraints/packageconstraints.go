package packageconstraints

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myNumber interface {
	~int | ~float64
}

func addT[T myNumber](a, b T) T {
	return a + b
}

type myAlisas int

func PackageExample() {
	// This function is intentionally left empty.
	// It serves as a placeholder to demonstrate the package structure.
	fmt.Println("Package constraints example")
	fmt.Println("Adding integers:", addI(3, 4))

	fmt.Println("Adding floats:", addF(3.5, 4.5))
	// Using the generic function with type constraints
	var n myAlisas = 5
	fmt.Println("Adding generic integers with myAlisas:", addT(n, 10))
	fmt.Println("Adding generic integers:", addT(5.1, 10.2))

}
