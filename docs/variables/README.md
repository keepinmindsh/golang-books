# 변수 

값을 저장하는 메모리 공간을 가리키는 이름.   

프로그램이란 결국 데이터를 연산/조작 하는 일.  

```go
package variables 

var a int = 100

```

##### 숫자타입 

- uint8, uint16, uint32, uint64
- int8, int16, int64
- float32, float64

- byte
- rune : 문자 1개 
- int 
- unit

##### 그외 타입 

- bool
- string
- 배열
- 슬라이스 
- 구조체 
- 포이터
- 함수타입
- 맵
- 인터페이스 
- 채널 

##### 여러가지 변수 선업법 

```go
package main

var name string = "Hello World" 

var name1 = "H"

name1 := "H"
```

##### 타입 변환 

- 연산의 각 항목 타입은 **반드시** 같아야한다. 