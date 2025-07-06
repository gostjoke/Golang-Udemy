package main

import (
	"fmt"

	"github.com/gostjoke/Golang-Udemy/Puppy"
)

func main() {
	Puppy.Bark()
	d := Puppy.Dog{Name: "Lucky"}
	d.Speak()
	fmt.Println("This is a simple Go program demonstrating package usage.")
}
