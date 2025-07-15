package callback

import "fmt"

func Printsquare(f func(int) int, i int) string {
	x := f(i)
	return fmt.Sprint("x squared is: ", x, " and the original number is: ", i)
}

func Square(x int) int {
	return x * x
}

func Callback() {
	fmt.Println(Printsquare(Square, 5))
}
