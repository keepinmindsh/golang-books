package main

import (
	"fmt"
	"golang-books/basic"
	protocol_client "golang-books/protocol/client"
	protocol_server "golang-books/protocol/server"
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

	// STRUCT
	var s basic.Student = basic.Student{
		ClassInfo: basic.ClassInfo{Class: 10, No: 10},
		Name:      "Seung Hwa",
		No:        1,
	}

	fmt.Print(s)

	go protocol_client.Server()

	protocol_server.Client()
}
