package slicefun

import (
	"sort"
	"strconv"
)

type Person struct {
	FirstName string
	Age       int
	Facvorite []string
}

func (p Person) Info() string {
	return p.FirstName + " is " + strconv.Itoa(p.Age) + " years old."
}

type ByAge []Person

func (a ByAge) Len() int      { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func SortPersonExp() {
	people := []Person{
		{FirstName: "Alice", Age: 30, Facvorite: []string{"Reading", "Traveling"}},
		{FirstName: "Bob", Age: 25, Facvorite: []string{"Gaming", "Cooking"}},
		{FirstName: "Charlie", Age: 35, Facvorite: []string{"Hiking", "Photography"}},
	}

	// Sort by age
	sort.Sort(ByAge(people))

	for _, person := range people {
		println(person.Info())
	}
}
