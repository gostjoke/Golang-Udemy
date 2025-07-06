package main

import "fmt"

func Bark() string {
	fmt.Println("Barking from the Puppy package")
	return "Woof!"
}

type Dog {
	Name string
}

unc (d Dog) Speak() {
	sound := Bark()
	fmt.Println(d.Name, "says:", sound)
}