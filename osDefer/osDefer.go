package osDefer

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func DeferExample() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
		panic(err)
	}
	entires, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
		panic(err)
	}
	for _, entry := range entires {
		fmt.Println("Entry Name: ", entry.Name())
	}
	// Defer will close the file after the function returns
	f, err := os.Open("osDefer\\test.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
		panic(err)
	}
	fmt.Println("File opened successfully")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
