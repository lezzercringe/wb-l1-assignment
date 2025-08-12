package main

import "fmt"

func set[T comparable](slice []T) []T {
	m := make(map[T]struct{}, len(slice))
	s := make([]T, 0, len(slice))

	for _, v := range slice {
		if _, ok := m[v]; ok {
			continue
		}

		m[v] = struct{}{}
		s = append(s, v)
	}

	return s
}

func main() {
	fmt.Println(
		set([]string{"cat", "cat", "dog", "cat", "tree"}),
	)
}
