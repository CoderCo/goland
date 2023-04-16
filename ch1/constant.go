package main

// Гломальные константы
const someInt = 1
const myInt int32 = 17
const myName = "coderco"

// Група констант
const (
	flag1 = 1
	flag2 = 2
)

// Група констант + iota
const (
	zeroVar = iota
	oneVar
	_        // пустая переменная, пропуск iota
	threeVar // 3
)

// арифметич.
const (
	z1 = iota + 10
	_
	z3
)

func main() {
	pi := 3.14

	// тип константы может быть определён во время компиляции
	println(pi + someInt)

	println(zeroVar, oneVar, threeVar, z3)
}
