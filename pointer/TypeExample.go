package pointer

import "fmt"

type Dog struct {
	name string
}

func (d Dog) Walk() {
	fmt.Println("Dog walking: ", d.name)
}

func (d *Dog) Run() {
	d.name = "Running " + d.name
	fmt.Println("Dog runing: ", d.name)
}

type Youngin interface {
	Walk()
	Run()
}

func YounginRun(y Youngin) {
	y.Run()
}

func DogExample() {
	d1 := Dog{name: "Buddy"}
	d1.Walk()                                   // Call method on value receiver
	d1.Run()                                    // Call method on pointer receiver
	fmt.Println("Dog name after run:", d1.name) // d1 got modified by pointer receiver

	d2 := &Dog{name: "Charlie"}
	YounginRun(d2) // Call method on pointer receiver through interface
}
