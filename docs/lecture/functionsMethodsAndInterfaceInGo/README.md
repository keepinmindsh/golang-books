# Functions

- 재사용성 
- 자주 사용되는 로직을 추출하여 하나의 메소드로 
- 함수는 1회 한번만 선언하면 된다 
- 함수 명으로 어떤 일을 하는지 유추할 수 있다. 
- 함수는 추상화의 목적으로 사용할 수 있다. 
  - 세부내용을 숨기는 역할을 할 수 있다. 
  - 단순히 프로세스만 이해하는 것으로 동작을 이해하면 된다. 어떻게 작성됬는지 뭐가 중요한가? 
  - 함수 추상화를 통해서 이해도를 높일 수 있고! 
  - 명확성을 위해서 명확한 이름이 중요하다! 

### Function Parameters 

- 동일 타입이 여러번 인자로 호출될 경우 

```go
package main

func sample(x, y int)  {
    
}
```

- 반환 값 

```go
package main

func sample(x, y int) int  {
    
	return x * y 
}
```

- 다중 반환 값

```go
package main

import "fmt"

func sample(x, y int) (int, int) {

	return x, y
}

func main() {
	x, y := sample(10, 10)

	fmt.Println(x, y)
}
```

### Call By Value, Reference

```go
package main

import "fmt"

func foo(y int) {
	y = y + 1
}

func main() {

	x := 2
	foo(x)
	fmt.Println(x) // result is 2 ( Call By Value ) 
}
```

- 데이터 캡슐화에 장점이 있으며 
  - 함수의 변수는 오직 함수 내에서만 변경된다. 
- 복사하는 시간 
  - 오브젝트가 클 경우에는 값을 복사하는데 시간이 소요된다.

```go
package main

import "fmt"

func foo(y *int) {
	*y = *y + 1
}

func main() {
	x := 2
	foo(&x)
	fmt.Println(x) // result is 3 ( Call By Reference ) 
}
```

- 복사 하는 시간이 빠른다. 
- 데이터 캡슐화에 취약하다. 주소값으로 전달되기 때문에 !! 

### Passing Arrays and Slices 

- 배열의 인자는 복사된다. 
- 배열은 사이즈가 클수 있기 때문에 문제가 될 수 있다.

```go
package main

import "fmt"

func foo(x [3]int) int {
	return x[0]
}

func main() {
	a := [3]int{1, 2, 3}

	fmt.Println(foo(a))

}

```

- 배열 포인터를 이용한 값 저달 

```go
package main

import "fmt"

func foo(x *[3]int) int {
	(*x)[0] = (*x)[0] + 1
}

func main() {
	a := [3]int{1, 2, 3}

	fmt.Println(foo(&a))
}

```

- slices 후에 배열을 전달하는 방법
  - 슬라이스는 배열에 대한 포인터를 가지고 있다. 
  - 배열을 전달하지 말고 슬라이스 인자를 전달해서 사용하라!! 

  
