package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	var a, b, c big.Int
	a.SetString("1"+strings.Repeat("0", 30), 10)
	b.SetString("1"+strings.Repeat("0", 31), 10)

	c.Div(&b, &a)
	fmt.Println("b / a =", &c)

	c.Add(&b, &a)
	fmt.Println("b + a =", &c)

	c.Mul(&b, &a)
	fmt.Println("b * a =", &c)

	c.Sub(&b, &a)
	fmt.Println("b - a =", &c)
}
