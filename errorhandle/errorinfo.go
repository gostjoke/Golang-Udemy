package errorhandle

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("norgate math error")

func ErrorInfoExample(n float64) {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(n)
	if err != nil {
		log.Fatalln("Error:", err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return f * f, nil
}
