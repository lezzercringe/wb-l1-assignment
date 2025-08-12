package main

import (
	"cmp"
	"fmt"
)

func find[T cmp.Ordered](s []T, target T) int {
	if len(s) == 0 {
		return -1
	}

	low, high := 0, len(s)-1

	for low <= high {
		i := (high + low) / 2

		if target == s[i] {
			return i
		}

		if target > s[i] {
			low = i + 1
			continue
		}

		if target < s[i] {
			high = i - 1
			continue
		}
	}

	return -1
}

func main() {
	fmt.Println(find([]int{1, 2, 10, 31}, 11))
}
