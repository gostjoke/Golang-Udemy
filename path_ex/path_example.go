package path_ex

import (
	"os"
)

func PathExample() {
	// This is an example function to demonstrate how to use the path_ex package.
	// You can add your own logic here.
	home := os.Getenv("OS")
	println(home)

	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(dir)
}
