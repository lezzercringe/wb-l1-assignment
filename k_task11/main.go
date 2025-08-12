package main

import "fmt"

func intersect[T comparable](a, b []T) []T {
	m := map[T]struct{}{}
	s := make([]T, 0)

	for _, v := range a {
		m[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := m[v]; ok {
			s = append(s, v)
		}
	}

	return s
}

func main() {
	fmt.Println(intersect(
		[]int{1, 2, 3},
		[]int{2, 3, 4},
	))
}
