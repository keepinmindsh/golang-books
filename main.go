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

	basic.TestPointer2()

	basic.Print()

	// STRUCT
	var s basic.Student = basic.Student{
		ClassInfo: basic.ClassInfo{Class: 10, No: 10},
		Name:      "Seung Hwa",
		No:        1,
	}

	fmt.Print(s)

	basic.ConstAndVar()

	// 반환되는 변수중 Align 하지 않으면셔 값을 사용하지 않는 방법 => _ 활용
	totalLength, _ := basic.LenAndUpper("Fighting!!")
	fmt.Print(totalLength)

	basic.AlignNothingForVariables()

	length, uppercase := basic.DeferSample("Fighting")
	fmt.Print(length, uppercase)

	//go protocol_client.Server()

	//protocol_server.Client()

	// github Library 모듈
	//ßßutil.HttpCall()

	basic.InterfaceSample()

}
