package tips

import "fmt"

// NotPointerFunc Go는 프로그램내에서 값의 참조와 레코드를 전달할 수 있는 pointers를 지원합니다.
func NotPointerFunc(ival int) {
	ival = 0
}

func PointerFunc(iptr *int) {
	/*
		그에 반해 zeroptr는 int형 포인터를 받도록 *int 타입을 파라미터로 갖고 있습니다.
		함수안에서 *iptr는 메모리 주소에서 해당 주소의 현재값으로 포인터를 역참조(dereference) 합니다.
		역참조된 포인터에 값을 할당하면 이는 참조된 주소의 값을 바꿉니다
	*/
	*iptr = 0
}

func TestPointer() {
	i := 1
	fmt.Println("initial:", i)
	NotPointerFunc(i)
	fmt.Println("zeroval:", i)

	PointerFunc(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

// TestPointer2 포인터의 역참조를 통해 메모리 주소로 value값 접근이 가능할 수 있었다.
func TestPointer2() {
	var example int
	var pointer *int
	example = 3

	pointer = &example

	fmt.Println(&example)

	// 포인터의 역참조 (*포인터변수명 = 대입값)
	*pointer = 5

	fmt.Println(pointer)

	fmt.Println(example)
}

// 포인터를 인자로 받는 함수
func EditNamePointer(item *SomeItem) {
	item.name = "호리병"
	fmt.Println(&item.name)
}

type SomeItem struct {
	id   int
	name string
}

// 구조체의 값을 그대로 인자로 받는 함수
func EditName(item SomeItem) {
	item.name = "도자기"
	fmt.Println(&item.name)
}

func sample() {

	item := SomeItem{1, "항아리"}
	fmt.Println("--- 객체이름 변수의 주소와 원래의 값 ---")
	fmt.Println(&item.name)
	fmt.Println(item.name)

	// 함수의 파라미터로 item 구조체를 전달
	fmt.Println("--- 객체이름 변경 함수의 주소와 예상 변경값 ---")
	EditName(item)
	fmt.Println(item.name)

	// 포인터를 함수의 파라미터로 전달
	fmt.Println("--- 객체이름 변경 함수의 주소와 예상 변경값(포인터 사용) ---")
	EditNamePointer(&item)
	fmt.Println(item.name)
}
