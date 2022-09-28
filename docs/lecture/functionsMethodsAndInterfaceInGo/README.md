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


### 함수를 잘 작성하는 법 

- 이해도가 높아야 합니다. 
  - 코드는 함수와 데이터
  - 함수는 함수자체로 이해 될 수 있도록 잘 표현되어야 함

- Debugging Principles
  - 에러의 2가지 요소 
    - 함수가 부정확하게 작성되었을 때 
    - 함수가 사용하는 데이터가 부정확할 때
- Supporting Debugging 
  - 함수는 이해할 수 있어야 한다. 
  - 데이터는 추적가능해야 한다. 
    - 전역 변수를 활용하는 방식은 좋지 않다! 


- 함수를 잘 작성할 수 있는 규칙 
  
1. 함수명 - 좋은 이름을 쓸 것!    

   모두가 이해할 수 있는 코드명/변수명을 사용해야합니다.

```go
package main 

// 명칭이 잘못된 경우
func processArray(a []int) float64 {
    return 1.23
}

// 명칭이 잘 작성된 경우 - 도메인이 특정된 이름으로 해당 도메인을 이해하고 있다면 도움이 됨. 
func computeRMS(samples []float64) float64 {
	return 1.23
}
```

2. 함수 응집도를 유지할 것 

"operation" 은 너무 많은 의미를 내포하고 있습니다. 응집도가 떨어집니다. 

```go
package main
// 좋은 예

func CountDigit()  {
  
}
```

3. 최소한의 파라미터 

- 너무 많은 파라미터는 함수 추적/파악을 어렵게 하고 
- 함수의 응집도를 떨어뜨린다. 
- 파라미터의 수를 최소화해야한다. 
  - 이를 위해서 파라미터를 하나의 구조체에 그룹지어서 전송해야할 필요도 있다.
  - 해당 구조체는 반드시 연관된 성향의 파라미터여야 한다. 

4. 함수 복잡도 

- 함수 길이는 복잡도의 기준이 됨. 
- 함수는 심플해야함. 
- 작은 함수도 복잡할 수 있음 

5. 함수 길이 

- 복잡한 프로세르르 어떻게 심플한 함수로 작성할 수 있는가? 
  - 함수 호출을 상속 구조로 작성한다. 
    - 함수하나가 다른 함수들을 구성하는 방식 

6. 제어 흐름 복잡도 

- 제어흐름은 조건에 따른 로직이 나뉘어 진다. 
- 조건을 분리한다. 
  - 조건에 따라서 함수로 분리하여 호출할 수 있도록 한다. 
    - 제어 흐름 상에 업무 로직의 노출이 아닌 함수의 노출이 중요하다. 


## Functions are First-class 

- Function은 다른 유형들처럼 다뤄진다! 
  - 동적으로 함수를 생성할수 있다. 
  - 변수는 함수 타입으로 선언될 수 있다. 
  - 인자로서 전달될 수 있으면서 값으로써 반환될 수 있다. 
  - 데이터 구조내에 저장될 수 있다.

```go
package main

import "fmt"

var funcVar func(int) int

func incFn(x int) int {
  return x + 1
}

func main() {
  funcVar = incFn

  fmt.Print(funcVar(1))
}
```

- 함수는 () 없이 함수의 오른쪽 편에 사용될 수 있다. 

### 인자로서의 함수 

- 함수는 다른 함수의 인자로서 전달 될 수 있음

```go
package main 

func applyIt(afunct func (int) int, val int) int {
	return afunct(val)
}
```

### 익명함수 

- 함수에 대한 이름을 지을 필요가 없음

```go
package main

import "fmt"

func applyIt(afunct func(int) int, val int) int {
  return afunct(val)
}

func main() {
  v := applyIt(func(x int) int { return x + 1 }, 2)
  fmt.Println(v)
}
```

## 리턴 값으로서의 함수

- 함수가 인자로써 전달 된다. 
- 함수가 반환될 함수로써 구성된다.

```go
package main

import (
  "fmt"
  "math"
)

func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
  fn := func(x, y float64) float64 {
    return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
  }

  return fn
}

func main() {
  Dist1 := MakeDistOrigin(0, 0)
  Dist2 := MakeDistOrigin(2, 2)
  fmt.Println(Dist1(2,2))
  fmt.Println(Dist2(2,2))
}
```

## 함수의 Scope

- 함수내에서 모든 명칭은 유효해야 한다. 
- 이름은 함수내에서 지역적으로 정의되어야 한다, 
- [Lexical Scope](https://www.techtarget.com/whatis/definition/lexical-scoping-static-scoping)
- 함수는 블록안에 정의된 이름을 포함한다.

```go
package main

import "fmt"

var x int

func foo(y int) {
  z := 1
  x = 10
  // x := 10 이코드는 에러가남...??? Why ? 

  fmt.Println(x, z)
}

func printX()  {
  fmt.Println(x)
}
```

### Closure 

- 함수 + 함수의 Scope
- 함수가 자신의 Scope 를 변수들과 함께 반환하거나 전달할 때 

## Variadic and Deffered

### Variable Argument Number 

- 함수는 많은 수의 인자를 가질 수 있다. 
- ... 를 사용하여 변수 인자의 수를 축약할 수 있다. 
- Slice 처럼 다뤄진다.

- slice를 다중인자를 받을 수 있는 함수로 전달 하할 수 있다. 
- 전달 할 때 vslice... 과 같이 변수 마지막에 ... 를 붙여야한다. 

```go
package main

import "fmt"

func getMax(vals ...int) int {
  maxV := -1
  for _, v := range vals {
    if v > maxV {
      maxV = v
    }
  }
  return maxV
}

func main() {
  fmt.Println(getMax(1, 3, 6, 4))
  vslice := []int{1, 3, 6, 4}
  fmt.Println(getMax(vslice...))
}
```

### Deffered Functio Calls 

- deferred는 함수가 완결되기 전까지는 호출을 지연할 수 있다! 
- 주로 cleanup 을 위한 프로세스를 위해서 사용된다.

```go
package main

import "fmt"

func main() {
  defer fmt.Println("Bye!")
  
  fmt.Println("Hello!")
}
```

### Deferred Call Arguments 

- 지연된 호출의 인자는 즉시 평가된다!
  - 즉 Deffered 호출은 호출 시점의 인자의 상태를 저장해서 이를 사용하는 것이지 이후 해당 연관된 인자가 변경된다고 해서 deferred 호출 시점에는 영향을 미치지 못한다. 

```go
package main

import "fmt"

func main() {
  i := 1
  defer fmt.Println(i + 1)
  i++
  fmt.Println("Hello!")
}

```

## Class and Encapsulation

- 데이터 필드의 집합과 함수들이 잘 정의된 Responsibility 에 의해서 공유된다.

> 예)  
>  Point Class 
>  클래스 의도 : Geometry를 위한 프로그램 
>  데이터 필드 : x, y 
>  함수 :   
>     DistToOrigin(), Quadrant()
>     AddXOffset(), AddYOffset()
>     SetX(), SetY() 

- Object 
  - 클래스의 객체 

- 캡슐화 
  - 데이터는 프로그래머로부터 보호될 수 있다. 
  - 데이터는 메소드를 사용함으로써 접급한 수 있다, 
  - 우리는 데이터를 보호하기 위해서 프로그래머를 믿어서는 안된다.

## Go는 "Class" 키워드가 없음 

- 대부분의 언어는 class 키워드를 가지고 있음
- 데이터 변수와 함수는 class 블록 내에서 정의되어 있다. 

### Data를 사용해 함수를 연계하는 방법 

- Method는 receiver type 이라는 것을 가진다. 
- . 에 의해서 메소드를 호출 할 수 있다.

```go
package main

import "fmt"

type MyInt int

func (mi MyInt) Double() int {
  return int(mi * 2)
}

func main() {
  v := MyInt(3)
  fmt.Println(v.Double())
}
```

> Tip)   
>  receiver type에 pointer를 붙여야 하는 경우에는   
> * receiver의 값을 변경하고자 할 때(단순히 읽기가 아닌 쓰기 작업도 같이)  
> * struct의 크기가 커서 deep copy 비용이 클 때  
> * 코드 일관성(option): 어떤 함수가 포인터 receiver를 사용하는 경우 일관성을 줄 때

```go
package main

import "fmt"

type Person struct {
  age int
}

func (p *Person) setAge(age int) {
  p.age = age
}

func (p Person) getAge() int {
  return p.age
}

func main() {
  p := Person{age: 3}
  p.setAge(10)
  fmt.Println(p.getAge)

  //OUTPUT
  //10
}

```

