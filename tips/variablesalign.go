package tips

import (
	"fmt"
	"strings"
)

// Const vs var
// 재할당이 불가능한 상수(const)와 재할당이 가능한 변수(var)로 나뉜다.
func ConstAndVar() {
	const name string = "unani"
	var number int = 123
	// 변수선언 축약 형태
	/*
		greeting := "Hello world" 와 같은 형태의 축약형도 사용이 가능하다.
		축약형은 상수가 아닌 변수로 생성이 가능하며 타입은 알아서 찾아서 매칭해준다.
		따라서 축약형으로 변수 선언 후 재할당 시 기존 타입과 다른 타입의 값으로 할당하면 에러가 발생한다.
	*/
	greeting := "Hello world"
	fmt.Println(name)
	fmt.Println(greeting)
	fmt.Println(number)
}

// Mutiple Values 반환하기
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func AlignNothingForVariables() {
	totalLength, _ := LenAndUpper("unani")
	fmt.Println(totalLength)
}
