package basic

import "fmt"

/*
	- zero value는 명시적인 초기값을 할당하지 않고 변수를 만들었을 때 해당 변수가 갖게 되는 값이다.
	- nil은 포인터, 인터페이스, 맵, 슬라이스, 채널, 함수의 zero value이다.
	출처: https://2kindsofcs.tistory.com/3 [세상에는 두 종류의 cs가 있습니다.:티스토리]
*/
func GoNilSample(args []string) bool {
	return args != nil
}

func GoStringCannotHaveNil() {
	var value string

	// 아래의 코드는 애시당초에 에러가 발생함, 문자열은 nil을 가질 수 없다.
	//fmt.Println(value == nil)

	// 이건 가능함.
	fmt.Println(value)
}

func GoNilWithObject() {
	var a []int
	b := make([]int, 0)
	c := []int{}
	var d []int
	d = nil
	e := make([]int, 0, 0)
	f := make([]int, 0, 1)

	fmt.Println(a == nil, // true
		b == nil, // false
		c == nil, // false
		d == nil, // true
		e == nil, // false
		f == nil) // false
}

func GoCannotCompareBetweenSlice() {
	var a []int
	//var b []int

	fmt.Println(a == nil)
	// 서로 다른 두 Slice 를 비용할 수 없다. 컴파일 에러 발생
	//fmt.Println(a == b)

	// todo : json serializing 하면 nil slice는 null로, empty slice는 []로 대응되었던 것 같다는 제보를 입수. 비어있는 것과 애초에 아무 것도 없는 것은 분명히 다르다.
}

func NilCompareClient() {
	var x *int32 = nil

	compare(x)

	var y *int32 = nil
	var z *int64 = nil
	compareBetweenNils(y, z)
}

func compare(x interface{}) {

	// nil은 변수의 데이터 타입도 저장하고 있기 때문에 비교를 하기 위해서는 비교하고 싶은 타입으로 형변환을 해서 사용해야 합니다.
	// Go에서는 nil을 사용합니다. 다른 언어와 다른 점으로는 Go에서는 nil은 데이터 타입 정보를 기억하고 있다는 것입니다.
	fmt.Println(x == (*int32)(nil))
}

func compareBetweenNils(x, y interface{}) {
	fmt.Println(x == y)
}
