# Go

오픈 소스 프로그래밍 언어임!

- 높은 생산성 
  - 간결하게 표현 가능한 코드 
  - 빠른 컴파일 속도 
  - 등 몇가지 이유 
- 쉬운 동시성 코드 작성 가능 

### Go 의 기본 형식 

- 구문 분석기를 통하여 문장 끝에 자동으로 붙는 세미콜론 
- Package main - import - func main() 순서 
- 변수 이름 -> 자료형 순서의 변수 선언 
- 다른 언어와 비교되는 조건문, 반복문의 형식 
  - 괄호가 필요 없음 
  - 타 언어의 while 문과 for문을 for 하나로 표현 가능 

### 기본 인코딩이 UTF-8임. 

##### 문자코드 : ANSI 

### Go Build

Go Build의 경우 module 없이는 빌드가 되지 않음.

go module를 생성해야함

```shell

go mod init goprojects/hello

```

##### 코드의 시작은 package 로 부터 반드시 시작해야한다. 주석은 제일 처음에 온다고 해서 문제 되지 않음.

##### 패키지는 자기자신을 import 할수 없음.

##### 주석

```go
// Package sample 
package sample

//

/* ~ */

```