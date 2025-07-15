package pointer

import "fmt"

type Person struct {
	name string
}

func newPerson(name string) *Person {
	p := Person{name: name} // 分配在 heap：因為回傳的是指標
	return &p
}

func localPerson(name string) Person {
	p := Person{name: name} // 分配在 stack：因為直接回傳值
	return p
}

func HeapOrStack() {
	p1 := newPerson("Alice")     // heap allocation
	p2 := localPerson("Charlie") // stack allocation

	fmt.Println(p1.name)
	fmt.Println(p2.name)
}
