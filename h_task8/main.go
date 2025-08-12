package main

import (
	"fmt"
	"os"
)

func btoi(value bool) int {
	if value {
		return 1
	}

	return 0
}

func setNthBit(num *int, ix int, value bool) {
	bitmask := 1 << ix
	*num = (*num &^ bitmask) | (-btoi(value) & bitmask)
}

func main() {
	i := 0
	setNthBit(&i, 0, true)
	setNthBit(&i, 2, true)
	if i != 5 {
		fmt.Println("FAILURE")
		os.Exit(1)
	}

	setNthBit(&i, 0, false)
	if i != 4 {
		fmt.Println("FAILURE")
		os.Exit(1)
	}

	fmt.Println("SUCCESS")
}
