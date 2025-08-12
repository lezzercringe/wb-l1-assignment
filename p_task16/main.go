package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"os"
	"slices"
)

func partition[T cmp.Ordered](s []T) int {
	pivot := s[0]
	i, j := 0, len(s)-1

	for ; ; i, j = i+1, j-1 {
		for i <= j && s[i] <= pivot {
			i++
		}

		for i <= j && s[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		s[i], s[j] = s[j], s[i]
	}

	s[0], s[j] = s[j], s[0]
	return j
}

func sort[T cmp.Ordered](s []T) {
	if len(s) <= 1 {
		return
	}

	sep := partition(s)

	sort(s[:sep])
	sort(s[sep+1:])
}

func Quicksort[T cmp.Ordered](slice []T) []T {
	c := slices.Clone(slice)
	sort(c)
	return c
}

func randomInts() []int {
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = rand.Int()
	}
	return arr
}

func main() {
	for range 100 {
		slice := randomInts()
		slice = Quicksort(slice)

		for i := range slice {
			if i == 0 {
				continue
			}

			if slice[i-1] > slice[i] {
				fmt.Print("NOT WORKING")
				os.Exit(1)
			} else {
				fmt.Println("OK")
			}
		}
	}
}
