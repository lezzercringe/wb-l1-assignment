package main

import "fmt"

func main() {
	a, b := 123, 1000
	b, a = a, b
	fmt.Println(a, b)

	// or

	a, b = 123, 1000
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)
}
