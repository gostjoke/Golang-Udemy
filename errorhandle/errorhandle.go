// =======================================
// @File        : errorhandle.go
// @Author      : Tien-Wei Hsu
// @Date        : 2025-07-20 22:41
// @LastEditTime: 2025-07-20 22:41
// @Description : Description
// =======================================

package errorhandle

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ErrorhandleExample() {
	var answer1, answer2, answer3 string

	fmt.Println("Name: ")
	_, err := fmt.Scanln(&answer1) // = input
	if err != nil {
		panic(err)
	}

	fmt.Println("Favorite Food: ")
	_, err = fmt.Scanln(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Favorite Sport: ")
	_, err = fmt.Scanln(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s, Favorite Food: %s, Favorite Sport: %s\n", answer1, answer2, answer3)
}

func ErrorhandleExampleDocument() {
	f, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		panic(err)
	}

	defer f.Close()

	r := strings.NewReader("Wassup\n") // Create a new reader with the string "Wassup"
	d := strings.NewReader("dog")      // Create a MultiWriter to write to both file and stdout
	io.Copy(f, r)
	io.Copy(f, d) // Write the content of the reader to the file and stdout
}
