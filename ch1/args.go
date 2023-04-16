// аргументы командной строки
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	var test int8 = 123
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		println(test)
		sep = " "
	}
	fmt.Println(s)
}
