package main

import "fmt"

type ByteSize int

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10 * 0) = 1 << 0 = 1
	MB                             // 1 << (10 * 1) = 1 << 10 = 1024
	GB                             // 1 << (10 * 2) = 1 << 20 = 1048576
	TB                             // 1 << (10 * 3) = 1 << 30 = 1073741824
	PB                             // 1 << (10 * 4) = 1 << 40 = 1099511627776
	EB                             // 1 << (10 * 5) = 1 << 50 =
	// ZB
	// YB
)

func MeasureBit() {
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)
	fmt.Printf("1 TB = %d bytes\n", TB)
	fmt.Printf("1 PB = %d bytes\n", PB)
	fmt.Printf("1 EB = %d bytes\n", EB)

	// Example of bitwise operations
	a := 5                                            // 0b0101
	b := 3                                            // 0b0011
	fmt.Printf("a & b = %d (binary: %b)\n", a&b, a&b) // Bitwise AND
	fmt.Printf("a | b = %d (binary: %b)\n", a|b, a|b) // Bitwise OR
	fmt.Printf("a ^ b = %d (binary: %b)\n", a^b, a^b) // Bitwise XOR

}

func main() {
	MeasureBit()
	fmt.Println("Measure Bit Example Completed")

}
