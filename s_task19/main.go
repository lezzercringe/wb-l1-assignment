package main

import (
	"fmt"
	"slices"
)

func reverse(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}

func main() {
	str := "главрыба ёё Ѵ"
	fmt.Println("Reversed", reverse(str))
	fmt.Println("Original", str)
}
