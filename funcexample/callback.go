package funcexample

func Callback() {
	x := doMath(2, 3, add)
	y := doMath(5, 2, subtract)
	println("Add result:", x)
	println("Subtract result:", y)
}

func doMath(x int, y int, c func(int, int) int) int {
	return c(x, y)
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}
