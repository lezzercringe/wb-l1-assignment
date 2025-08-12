package main

import (
	"fmt"
	"runtime"
)

func remove[T any](s []T, i int) []T {
	copy(s[i:], s[i+1:])
	// в результате предыдущей операции, последний элемент слайса в памяти которую референсит слайс задублируется
	//
	// в случае если это был указатель/slice стоит обнулить его предыдущую позицию (которая окажется вне границ слайса),
	// чтобы GC смог собрать референсируемое им значение, когда он будет удален со своей новой позиции в слайсе
	var zero T
	s[len(s)-1] = zero
	return s[:len(s)-1]
}

func main() {
	mustBeFreed := make([]byte, 1<<10)
	runtime.SetFinalizer(&mustBeFreed, func(s *[]byte) {
		fmt.Println("FREED!")
	})

	s := [][]byte{{'a'}, {'b'}, {'X'}, {'d'}, mustBeFreed}
	mustBeFreed = nil // вообще работает и без этого, но для чистоты эксперимента стоит явно сделать nil

	fmt.Println("Initial:", s)
	s = remove(s, 2)
	fmt.Println("After removing X:", s)
	s = remove(s, 3)
	runtime.GC()
	fmt.Println("(^ FREED! must be up here ^) After removing BIG SLICE :", s)
}
