# Go

오픈 소스 프로그래밍 언어임! 

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