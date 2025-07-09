package main

import (
	"fmt"

	"github.com/gostjoke/Golang-Udemy/Iota"
	"github.com/gostjoke/Golang-Udemy/Puppy"
	"github.com/gostjoke/Golang-Udemy/ValueType"
	"github.com/gostjoke/Golang-Udemy/path_ex"
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
	path_ex.Godotenvtest()

}
