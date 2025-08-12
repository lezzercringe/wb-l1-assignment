package main

import (
	"fmt"
	"slices"
)

func main() {
	m := map[int][]float64{}
	buckets := make([]int, 0)
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, t := range temps {
		d := int(t) / 10

		if len(m[d]) == 0 {
			buckets = append(buckets, d)
		}

		m[d] = append(m[d], t)
	}

	slices.Sort(buckets)
	for _, bkt := range buckets {
		fmt.Printf("%v:%v\n", bkt*10, m[bkt])
	}
}
