package slicefun

import (
	"fmt"
	"sort"
)

func SortInt(T []int) {
	// Sort the slice in ascending order
	sort.Ints(T)
	fmt.Println("Sorted slice:", T)
}

func SortString(T []string) {
	// Sort the slice in ascending order
	sort.Strings(T)
	fmt.Println("Sorted slice:", T)
}

func SortEX() {
	a := []int{5, 3, 8, 1, 2}
	b := []string{"banana", "apple", "cherry", "date"}
	SortInt(a)
	SortString(b)
}
