package main

import (
	"fmt"
	"reflect"
)

func printTypename(v any) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		if reflect.ValueOf(v).Kind() == reflect.Chan {
			fmt.Println("chan")
		} else {
			fmt.Println("unmatched :(")
		}
	}
}

func main() {
	printTypename(make(chan int)) // chan
	printTypename(uint64(123))    // unmatched
}
