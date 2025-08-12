package main

import (
	"fmt"
	"task24/geometry"
)

func main() {
	a := geometry.NewPoint(10, 10)
	b := geometry.NewPoint(5, 30)

	fmt.Println(a.Distance(b))
}
