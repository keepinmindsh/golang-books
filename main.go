package main

import (
	"fmt"
	"golang-books/basic"
)

func main() {
	fmt.Println("Hello World!")

	basic.GoStringCannotHaveNil()

	basic.GoNilWithObject()

	a, b := basic.Swap("Hello", "World")
	fmt.Println(a, b)

	x, y := basic.Split(50)
	fmt.Println(x, y)

	basic.ValueCheck()

	basic.TestPointer()

	basic.Print()
}
