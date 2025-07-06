package Puppy

import (
	"fmt"
	// "time"
)

type Person struct {
	Name       string
	Age        int
	City       string
	Occupation string
}

func main() {
	person := Person{Name: "Bob", Age: 25, City: "Seattle", Occupation: "Designer"}
	fmt.Printf("Person Details: %+v\n", person)
}
