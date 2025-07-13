package structexp

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Favorite  []string
	Friends   []Person
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) Speak() string {
	return fmt.Sprintf("Hello, my name is %s %s and I am %d years old.", p.FirstName, p.LastName, p.Age)
}

func NewMap() {
	p1 := Person{
		FirstName: "Alice",
		LastName:  "Smith",
		Age:       28,
		Favorite:  []string{"Reading", "Hiking"},
	}
	p2 := Person{
		FirstName: "Bob",
		LastName:  "Johnson",
		Age:       35,
		Favorite:  []string{"Cooking", "Traveling"},
		Friends:   []Person{p1},
	}
	m := map[string]Person{
		p1.LastName: p1,
		p2.LastName: p2,
	}
	fmt.Println("Map of Persons: ", m)
}
