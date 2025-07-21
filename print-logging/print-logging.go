package printlogging

import (
	"fmt"
	"log"
	"os"
)

func PrintLoggingExample() {
	defer foo()
	_, err := os.Open("example.txt")
	fmt.Println("This is a print logging example")
	if err != nil {
		log.Fatalln(err)
		// return
	}
	// if
}

func foo() {
	fmt.Println("This is a deferred function")
}

func CreateLogFile() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer f.Close()

	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")

	if err != nil {
		log.Println("Error:", err)
	}
	defer f2.Close()

	fmt.Println("This will not be logged to the file")

}
