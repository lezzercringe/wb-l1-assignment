package main

import "fmt"

func main() {
	source := make(chan int)
	double := make(chan int)

	nums := []int{1, 2, 3, 4, 10}

	// source writer
	go func() {
		for _, x := range nums {
			source <- x
		}
		close(source)
	}()

	// double writer
	go func() {
		defer close(double)
		for x := range source {
			double <- x * 2
		}
	}()

	// checker
	i := 0
	for d := range double {
		if nums[i]*2 != d {
			fmt.Println("wrong result")
		} else {
			fmt.Println("ok")
		}

		i++
	}

}
