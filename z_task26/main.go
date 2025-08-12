package main

import (
	"fmt"
	"strings"
)

func areAllUnique(str string) bool {
	m := map[rune]struct{}{}

	for _, c := range strings.ToLower(str) {
		if _, ok := m[c]; ok {
			return false
		}

		m[c] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(
		areAllUnique("abcd"),
		areAllUnique("abCdefAaf"),
		areAllUnique("aabcd"),
	)
}
