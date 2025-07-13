package interfacespolymorphstringer

import (
	"fmt"
	"log"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprintf("Book Title: %s", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprintf("Count: %d", c)
}

func logInfo(s fmt.Stringer) {
	log.Printf("Logging info: %s", s.String())
}

func BookExample() {
	b := book{title: "The Go Programming Language"}
	c := count(42)

	logInfo(b)
	logInfo(c)
}
