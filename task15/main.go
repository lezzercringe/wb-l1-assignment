package main

import "strings"

// В реализации из условия полная строка не может быть собрана сборщиком мусора из-за того,
// что на нее ссылается justString (глобальная переменная), по-скольку slice референсит изначально аллоцированный кусок памяти.
// Для исправления ситуации достаточно хранить префикс не слайсом большой строки, а копией точно нужного размера.

func createHugeString(n int) string

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}
