package main

import "fmt"

// Глобальная константа
const size uint = 5

func main() {
	var a1 [5]int
	fmt.Println("массив", a1, "длина", len(a1))

	// Устанавливаем размер масива используя константу
	var a2 [5 * size]bool
	fmt.Println(a2, "длина", len(a2))

	// Устанавливается длина при инициализации
	a3 := [...]int{3, 3, 3, 4, 5}
	fmt.Println("длина при инициализации", a3, "длина", len(a3))

	// Двухмерный массив
	var aa [3][3]int
	aa[1][1] = 1
	fmt.Println("двухмерный массив", aa)
}
