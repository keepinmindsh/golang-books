package variables

import "fmt"

// 변수 - 값을 저장하는 메모리 공간을 가리키는 이름

// https://www.youtube.com/watch?v=-oqy4PAxC1k&list=PLy-g2fnSzUTBHwuXkWQ834QHDZwLx6v6j&index=7

func Variables() {

	var a int = 10
	var msg string = "Hello Varaible"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)
}