package main

import (
	"fmt"
	"slices"
	"unicode"
)

func reverseWordsOrder(s []rune) {
	slices.Reverse(s)

	l := 0
	for r, c := range s {
		if unicode.IsSpace(c) {
			slices.Reverse(s[l:r])
			l = r + 1
		}
	}

	slices.Reverse(s[l:])
}

func main() {
	str := []rune("snow dog sun")
	reverseWordsOrder(str)
	fmt.Println(string(str))
}
