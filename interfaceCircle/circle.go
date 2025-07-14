package interfaceCircle

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Shape interface {
	Area() float64
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func ExampleCircle(R float64) {
	c := Circle{
		Radius: R,
	}
	area := GetArea(c)

	fmt.Printf("Area of Circle: %v", area)
}
