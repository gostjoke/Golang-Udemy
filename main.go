package main

import (
	"fmt"

	"github.com/gostjoke/Golang-Udemy/Iota"
	"github.com/gostjoke/Golang-Udemy/Puppy"
	"github.com/gostjoke/Golang-Udemy/RaceCondition"
	"github.com/gostjoke/Golang-Udemy/ValueType"
)

func PuppyBark() {
	Puppy.Bark()
	d := Puppy.Dog{Name: "Lucky"}
	d.Speak()
}

func IotaExample() {
	Iota.Iota()
	ValueType.ValueTypeF()
}

func ChannelExample() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) // 關閉 channel 才能 for range
	}()

	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	// d := dictexample.Dictionary{}
	// d.Add("hello", "world")

	// controlflow.SwitchStatementExample()

	// channel.ChannelStateExample()
	// path_ex.PathExample()
	// path_ex.Godotenvtest()
	// container introduce
	// container.SliceArray()
	// container.HashmapExample()
	// a := []int{1, 2, 3, 4, 5, 6}
	// b := append(a[3:], a[1:]...)
	// fmt.Println("Original slice:", b)
	// slicefun.Reverse(b)
	// fmt.Println("Modified slice:", b)
	// slicefun.Makefun()
	// x := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	// fmt.Println("Median:", slicefun.Median(x))
	// structexp.NewMap()
	// structexp.GetCar()
	// Mycars := structexp.CreateMultCar(structexp.Engine{
	// 	EngineName: "V8",
	// 	EngineType: "Petrol",
	// 	HorsePower: 450,
	// }, "Car1", "Car2", "Car3")

	// // fmt.Println("Created Cars:", Mycars)
	// for _, car := range Mycars {
	// 	fmt.Println(car.Display())
	// }
	// osDefer.DeferExample()
	// interfacePolymorphism.SpeechExp()

	// interfacespolymorphstringer.BookExample()

	// interfacespolymorphstringer.WriterExample()
	// interfaceCircle.ExampleCircle(12.12)
	// funcexample.ExampleFunction()
	// funcexample.Callback()
	// funcexample.Closure()
	// funcexample.Fibonacci(0, 1, 100)
	// callback.Callback()
	// a := 3.0
	// pointer.Power2(&a)
	// fmt.Println("Power of 2:", a)
	// b := 12.0
	// pointer.Power3(&b)
	// fmt.Println("Power of 3:", b)
	// pointer.Dereferencing()
	// pointer.DogExample()
	// packageconstraints.PackageExample()
	// JsonMarshal.MarshalExpample()
	// writerinterface.WriteToConsole("Hello, World!\n")
	// Sort Slice
	// slicefun.SortPersonExp()
	// Hash Example
	// bcryptex.BcryptExample()
	// concurrency.WaitGroupExample()
	// concurrency.DocumentationExample()
	RaceCondition.AtomicExample()

}
