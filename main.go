package main

import (
	"fmt"
	"github.com/keepinmindsh/go-lang-module/util"
	"golang-books/problems"
	"golang-books/setting"
	"golang-books/tips"
)

func main() {

	//hello.HelloWorld()

	problems.PrintSong(20)

	problems.PrintHanoi(100)

	problems.Fibonacci(5)

	setting.Hanoi(3)

	util.Fib(10)

	//study001()
}

func study001() {
	fmt.Println("Hello World!")

	tips.GoStringCannotHaveNil()

	tips.GoNilWithObject()

	a, b := tips.Swap("Hello", "World")
	fmt.Println(a, b)

	x, y := tips.Split(50)
	fmt.Println(x, y)

	tips.ValueCheck()

	tips.TestPointer()

	tips.TestPointer2()

	tips.Print()

	// STRUCT
	var s tips.Student = tips.Student{
		ClassInfo: tips.ClassInfo{Class: 10, No: 10},
		Name:      "Seung Hwa",
		No:        1,
	}

	fmt.Print(s)

	tips.ConstAndVar()

	// 반환되는 변수중 Align 하지 않으면셔 값을 사용하지 않는 방법 => _ 활용
	totalLength, _ := tips.LenAndUpper("Fighting!!")
	fmt.Print(totalLength)

	tips.AlignNothingForVariables()

	length, uppercase := tips.DeferSample("Fighting")
	fmt.Print(length, uppercase)

	//go protocol_client.Server()

	//protocol_server.Client()

	// github Library 모듈
	//ßßutil.HttpCall()

	tips.InterfaceSample()
}
