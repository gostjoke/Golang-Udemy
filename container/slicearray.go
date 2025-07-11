package container

import (
	"fmt"
)

func SliceArray() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := []string{"a", "b", "c", "d", "e"}
	array := [5]int{1, 2, 3, 4, 5}
	array2 := [5]string{"a", "b", "c", "d", "e"}

	// array change to slice
	arraySlice := array[:] // 將 array 轉換為 slice
	fmt.Println("Array to Slice:", arraySlice)

	fmt.Println("Slice:", slice)
	fmt.Println("Slice2:", slice2)
	fmt.Println("Array:", array)
	fmt.Println("Array2:", array2)

	slice = append(slice, 6)     // 添加元素到 slice
	slice2 = append(slice2, "f") // 添加元素到 slice2
	fmt.Println("Updated Slice:", slice)
	fmt.Println("Updated Slice2:", slice2)

	// [...]int{7, 8, 9} = [3]int{7, 8, 9} // auto create

}
