package main

import (
	"fmt"

	"github.com/gostjoke/Golang-Udemy/Iota"
	"github.com/gostjoke/Golang-Udemy/Puppy"
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

}
