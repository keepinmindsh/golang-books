# Go Study 

### Package Import 

- Package 를 import 했을 때, 외부에 공개 가능한 함수는 영문자로 첫글자가 반드시 대문자여야한다. 

### Go - Pointer 를 쓰는 이유

- Stack & Heap 
  - 실행 시 동적으로 메모리를 확보하는 영역으로 스택과 힙이 있다. 
  - 스택 메모리는 함수 호출 스택을 저장하고 로컬 변수, 인수, 반환 값도 여기에 둔다.
  - 스택의 Push와 Pop은 고속이므로 객체를 스택 메모리에 저장하는 비용은 작다. 단 함수를 나오면 스택이 Pop 되어 해제되므로 함수의 수명을 넘는 객체는 살 수 없다.
  - 힙 메모리는 콜 스택과는 관계 없으므로 함수 범위에 얽매이지 않고 객체를 저장해 둔다. 다만 빈 영역을 찾고, GC로 쓸모 없게된 객체를 회수하기도 하므로 처리 비용이 든다.
  - Go 언어는 컴파일러가 객체를 스택에 확보할지 힙에 확보할지 결정하므로 프로그래머가 의식할 필요는 보통 없다.

- 이를 정리해보면 
  - 함수 내에서만 사용되는 값은 스택에 둔다.
  - 어떤 함수 내에서 확보한 값이 함수 밖에서도 필요하게 된다면 힙에 놓인다

##### Compiler 에 Flag를 전달하는 방법 

```shell
> go build -gcflags -m hello.go
```

##### Stack 

```go

func test1() {
    var d Duck = Duck{}
    d.Sound()
}

func test2() {
    var d *Duck = &Duck{}
    d.Sound()
}

func test3() Duck {
    var d Duck = Duck{}
    return d
}

func test4() Duck {
    var d *Duck = &Duck{}
    return *d
}

```

##### Heap

```go

func test5() *Duck {
    var d Duck = Duck{}
    return &d
}

func test8() Sounder {
  var d Sounder = &Duck{}
  return d
}

```

###### 정리

- 함수 내에서만 사용되는 값은 스택
- 함수 밖에서도 값이 필요하게 된다면 힙
- new 해도 스택인 경우가 있다
- 함수 내에서만 사용되는 값이면
- 로컬 변수의 주소를 반환해도 된다
- 힙에 두도록 컴파일러가 다룬다
- 인라인 전개에 따라 잘 최적화해 준다
- interface에 대입하면 힙

> [참조 문서 - Stack/Heap](https://jacking75.github.io/go_stackheap/)

###### 포인터를 쓰는 이유 

포인터는 타입이다. 대신 포인터가 가리키는 변수의 메모리 주소를 갖는다. 즉 포인터 변수를 통해 다른 변수의 메모리 주소를 참조해 무언가 가능할 것 같은 느낌이다.  

- 메모리 주소
- 역참조 

```go

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

```

go언어는 기본적으로 값에 의한 호출(Call by value)이므로 매개변수를 복사해서 함수 내부로 전달하는데 포인터를 통해 본래의 값을 변경할 수 있다.

### go 에서의 인터페이스 

go 에서 인터페이스는 Java와 달리 구현의 개념이 아닌 Interface가 실제 함수의 추상을 가지고 이를 사용할 수 있는 구조로 되어 있음 

```go


func (p Pet) getName() {
	fmt.Printf("pet name is %s.\n", p.Name)
}

func (p Pet) getAge() {
	fmt.Printf("pet age is %d.\n", p.Age)
}

type GetName interface {
	getName()
}

type GetAge interface {
	getAge()
}

type PetInfo interface {
	GetName
	GetAge
}

func interfaceSample() {
	var interfaceSample PetInfo

	interfaceSample.getName()
}


```

### Go Method Receiver 

Go는 클래스가 없다. 하지만 특정 타입에 대한 메소드를 만들 수 있다.  

Vertex와 Vertex3D 구조체가 있다.  

- Vertex는 Abs라는 메소드를 갖는다.

```go
package main

import "math"

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
```
- Vertex3D는 Abs 라는 함수를 통해 연산이 가능하다.

Go에서는 타언어의 인스턴스와 비교해보면 필요한 구조체만 해당 함수에 접근하면 되니까 메모리 다이어트가 필요 없게 되는 것 같다. 

##### Pointer Receiver 

- 메소드가 receiver pointer의 값을 수정할 수 있습니다. 
- 메소드 호출에 따른 값의 복사를 방지하기 위해서입니다. 구조체가 클 수록 효율이 좋습나다. 
