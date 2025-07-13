package interfacePolymorphism

import (
	"fmt"

	"github.com/gostjoke/Golang-Udemy/structexp"
)

type Human interface {
	Speak() string
}

func PrintSpeech(h Human) Human {
	fmt.Println(h.Speak())
	return h
}

func SpeechExp() {
	p := structexp.Person{
		FirstName: "Alice",
		LastName:  "Smith",
		Age:       28,
	}

	human1 := PrintSpeech(p)
	fmt.Println("Human1:", human1.Speak())

}
