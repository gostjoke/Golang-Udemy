package funcexample

func Closure() {
	x := incrementor()
	// x()
	println("Closure result:", x())
	println("Closure result:", x())
	println("Closure result:", x())
	println("Closure result:", x())
	println("Closure result:", x())
	println("Closure result:", x())

	g := incrementor()
	println("Closure result:", g())
	println("Closure result:", g())
	println("Closure result:", g())
	println("Closure result:", g())
	println("Closure result:", g())
	println("Closure result:", g())

}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
