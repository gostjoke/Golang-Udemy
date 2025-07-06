package Puppy

import "fmt"

func Bark() string {
	fmt.Println("Barking from the Puppy package")
	return "Woof!"
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	sound := Bark()
	fmt.Println(d.Name, "says:", sound)
}
