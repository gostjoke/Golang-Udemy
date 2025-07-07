package main

import (
	"fmt"

	controlflow "github.com/gostjoke/Golang-Udemy/ControlFlow"
	"github.com/gostjoke/Golang-Udemy/Iota"
	"github.com/gostjoke/Golang-Udemy/Puppy"
	"github.com/gostjoke/Golang-Udemy/ValueType"
	"github.com/gostjoke/Golang-Udemy/dictexample"
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
	d := dictexample.Dictionary{}
	d.Add("hello", "world")

	controlflow.SwitchStatementExample()
}
