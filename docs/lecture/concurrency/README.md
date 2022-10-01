## Parallel Execution 

- 2개의 프로그램에 완전히 동시에 진행되는 것을 예로 들수 있다. 
- P1, P2가 특정 시간 t에 동시에 시작되는 것을 말한다. 
  - CPU1, CPU2 에 의해서 동시에 실행되며 멀티 Core에서 실행할 수 있을 이야기 한다. 

## Why Use Parallel Execution

- 예) 2개의 접시를 동시에 씻을 수 있기 때문에 빠르다. 
- 일부의 업무는 순서가 반드시 존재한다. 
  - 이를 우리는 프로그램에 의해서 병렬적으로 시퀀셜하게 작성할 수 있다! 
  - 또한 아무래 병렬적으로 CPU 또는 자원이 제공된다하더라도 시퀀셜하게 가는 업무의 경우, 자원을 모두 사용하지 못할 수 있음 
  - 병렬적이라고 무조건 빠르게 처리할 수 있다는 것은 오산이다. 

## Why use Concurrency 

- 병렬성 없이 속도를 빠르게 하기 위해서 
  - Processor를 빠르게 만드는 것! 
  - 속도를 빠르게 하기 위해서 더 많은 메모리와 동작하도록 작업되었다. 
    - 폰 노이만 병목 현상을 최소화 하는 ㄱ4ㅓㅅ 

## Moore's Law



## Power/Temperature Problem 

- 트랜지스터는 그들이 전환 될 때 파워를 소모한다.
- 트랜지스터의 밀도는 높이는 것은 증가된 파워의 소비를 야기한다. 
- 높은 파워는 높은 온도를 야기한다. 
  - Air Cooling은 너무 많은 열을 식히기 위해서 동작한다. 

## Dynamic Power 

- P = a * CFV^2
- a is percent of time switching 
- C is capacitance ( related to size )
- F is the clock frequency
- V is voltage swing ( from low to high )
- Voltage is important 
- 0 to 5V는 0 to 1.3V 보다 더 많은 파워를 사용해야 한다.

## Dennard Scaling 

- 트랜지스터의 사이즈에 따라서 전압은 확장되어야 한다. 
- 파워 소비 및 온도를 유지해야한다. 
  - 전압은 너무 낮을 수 없다. 
  - 파워 손실은 고려하지 않는다. 

## MultiCore System 

- 빈도를 증가시킬 수 없음 
- 빈도 증가 없이 프로세스 코어를 추가할 수 있다. 
- 병렬 실행은 폭발적인 멀티 코어 시스템을 요구한다. 

> Tip) 
> 동시성 프로그래밍과 병렬성 프로그래밍의 차이  
> 동시성 프로그래밍과 병렬성 프로그래밍 모두 비동기 동작을 구현할 수 있지만 그 동작 원리가 다르다.  
>  * 동시성 : 싱글 코어, 멀티 코어에서 모두 구현 가능.  
> ex. 하나의 커피머신에 커피를 받기 위한 N개의 대기열 -> 서로 번갈아 가며 커피를 받아감.    
>  * 병렬성 : 멀티 코어에서만 구현 가능.  
> ex. N대의 커피머신에 커피를 받기 위한 N개의 대기열 -> 각 커피머신마다 하나의 줄을 가지고 있어 각각의 줄마다 커피를 받아감.  

## Concurrent Execution 

- 동시성은 완벽히 동시에 실행된다고 보장할 수 없다. 
  - 동시성은 하나의 동일한 하드웨어에서 실행된다고 할 수 있다.
    - 오직 하나의 Task 가 특정시간에 실행될 것이다. 
- 병렬성은 정확히 동일한 시간에 처리된다. 
  - 병렬성은 반드시 서로 다른 하드웨어에서 실행된다. 

두개의 Task 를 동시성과 병렬성을 차이를 두고 처리한다고 할 때,  

동시성은 하나의 Thread 에서 2개의 업무를 번갈아 수행할 것이며,   
병렬성은 2개의 Thread 에서 2개의 업무를 동시에 수행할 것이다. 

## Concurrent Programming

- 프로그래머는 어떤 Task 등을 병렬적으로 실행할 수 있는지 정한다. 

## Hiding Latency 

- 동시성은 성능을 증진시킬 수 있다. 
- Task 들은 반드시 주기적으로 기다린다. 
  - Wait for Memory 
  - X = Y + Z read Y, Z from memory 
  - 동시성 프로그래밍의 Task 들은 하나의 Task가 Memory를 읽거나, IO에 접근하는 대기 시간동안에 해당 Task외의 다른 Task가 업무를 볼 수 있다. 

## Hardware Mapping In Go 

- 프로그래머는 하드웨어 매핑을 결정하지 않는다. 
- 프로그래머는 병렬성을 가능하기 만든다. 
- 하드웨어 매핑은 다양한 요소에 영향을 받는다. 

## Concurrency Basic 

### Processes

### Memory

 - 가상의 주소 공간 
 - Code, Stack, Heap, Shared Library

### Registers 

 - Program Counter, Data Regs

### Operating Systems

- 많은 프로세스가 동시에 실행 될 수 있도록 허용한다.
- Process 간의 변경은 빠르게 일어난다. 
  - 20 ms 이내 

### Scheduling Process 

- 병력적인 실행의 환상을 심어주는 역할 
- Operating System은 실행을 위해서 프로세스의 실행 계획을 관리한다. 
- OS는 CPU, Memory 등등에 공평한 접근을 제공한다. 

### Context Switching 

- 제어흐름은 하나의 프로세스에서 다른 프로세르로 변경된다.
  - 이때 발생되는 것을 Context Switching이라고 한다. 

- 프로세스의 컨텍스트(문맥)은 반드시 변경되어 사용되어야 한다. 

### Thread And GoRoutines 

- Thread 는 Context를 공유한다. 
- 많은 스레드가 하나의 프로세스 안에 존재할 수 있다. 

### Goroutines 

- Go에서의 스레드와 같은 역할 
- 많은 고루틴은 Single OS Thread 안애서 실행된다.

### Go Runtime Scheduler 

- OS Thread 안에서 goroutines의 실행 계획을 관리한다. 
- 마치 작은 OS 처럼 동작한다. 
- Logical Processor 는 하나의 Thread 와 연결된다. 

### Interleavings 

- 하나의 Task 내에서 실행의 순서는 알수 있음. 
- 동시에 실행되는 Task 들 사이의 실행의 순서는 알수 없음. 
- Task들 사이의 Interleaving은 알 수 없음. 

- 많은 Interleavings이 가능함. 
  - 반드시 모든 가능성에 대해서 고려해야한다!
- 순서를 만드는 것은 결정가능한 항목이 아님 

### Race Conditions 

- 결과는 무질서한 순서에 의존한다. 
- 경쟁은 객체 간의 통신에 의해서 발생할 수 있다. 

## Goroutine

- One Goroutine 은 main() 함수를 실행할 때 자동으로 생성된다. 
- 다른 Goroutine 은 go keyword를 이용해서 생성할 수 있다.

아래의 코드에서 foo() 함수 뒤의 a = 2는 foo()가 끝나기 전까지는 실행될 수 없다. 

```go
package main

import "fmt"

func main() {
  a := 1
  foo()
  a = 2

  fmt.Println(a)
}

func foo() {
	fmt.Println("This is Foo!")
}
```

아래의 코드에서 foo() 함수 뒤의 a = 2는 foo() 끝나는 것을 기다리지 않는다.   
그 이유는 go foo() 라인이 실행될 때 새로운 GoRoutine이 생성되며 Go Runtime Schedule에 의해서 실행되지만  
이게 언제끝날지도 우리는 알수 없지만 a = 2는 foo() 함수가 끝나는 것을 기다리지 않는 상태로 완결된다는 것을 인식해야한다.   

```go
package main

import "fmt"

func main() {
  a := 1
  go foo()
  a = 2

  fmt.Println(a)
}

func foo() {
  fmt.Println("This is concurrent Foo")
}

```

### Existing a Goroutine 

- A GoRoutine는 Code가 완료되면 종료된다, 
- 만약에 Main Goroutine이 끝난다면, 그외의 다른 GoRoutine은 종료된다. 
  - 즉 만약 Main GoRoutine이 끝날 때, Other Goroutine의 실행 여부와 상관 없이 강제로 모두 종료시킨다. 

### Early Exit

아래의 코드에서 print의 우선순위는 우리가 정할 수 없다. 다만 Go Runtime Schedule은 Main Goroutine이 먼저 실행되는 것을 선호한다.  
**다만 이것은 Go 버전에 따라 달라질수 도 있는 일이다!** 

- 오직 "Main Routine" 만 출력된다. 
- 새로운 GoRoutine이 시작되기 전에 메인이 종료된다. 

```go
package main

import "fmt"

func main() {
  go fmt.Println("New Routine")
  fmt.Println("Main Routine")
}

```

### Delayed Exit

GoRoutine을 위해서 지연을 넣는 코드는 좋은 코드가 아님.  
아래와 같은 코드는 동작은 함... 하지만 저렇게 하면 의미가 없음... 

아래의 코드는 때때로 동작할수도 있고 때때로 에러를 일으킬 수 도 있다.

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  go fmt.Println("New Routine")
  time.Sleep(100 * time.Millisecond)
  fmt.Println("Main Routine")
}

```

### Synchronization 

전역 이벤트를 사용하는 것은 모든 스레드에 의해서 동시에 보여질 수 있는 위치에 있는 것이다. (동시에 접근할 수 있다는 것) 

- 모든 스레드에 하나의 값에 대해서 전역적으로 접근할 때, 해당 값을 접근하는 순서, 위치에 따라서 값이 변경될 수 있음 
  - go routine을 사용할 때 고려해야할 사항임. 
  - 우리가 원하는 오퍼레이션 시점에 값을 정의하고 싶을 때 Synchronization 을 관리할 수 있다. 
- 동기화는 성능을 느리게할 수 있지만 필수적으로 사용해야하는 경우가 발생할 수 있다. 

### Sync WaitGroup 

- Sync Package 는 goroutine 사이의 동기화를 하기위한 함수를 제공한다. 
- sync.WaitGroup는 GoRoutine에게 다른 GoRoutine을 강제적으로 기다리라고 하기 위한 것임!
- 고루틴이 생길 때마다 internal counter가 생성된다. 
  - sync.WaitGroup은 internal counter를 이용해서 고루틴을 대기 시킨다.

![Group Routing Sync Wait](https://github.com/keepinmindsh/golang-books/blob/main/docs/assets/Concurrency_005.png)

- Add() 는 숫자를 증가시킨다. 
- Done() 은 숫자를 감소시킨다. 
- Wait()는 숫자가 0이 될 때까지 대기한다.

```go
package main

import (
  "fmt"
  "sync"
)

func foo(wg *sync.WaitGroup) {
  fmt.Printf("New Routine")
    wg.Done()
}

func main()  {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	fmt.Printf("Main Routine")
}
```

## Goroutine Communication 

- Goroutine 는 하나의 큰 작업을 실행하기 위해서 함께 작업하는 경우가 많다. 
- 협력을 위한 데이터를 전송해야 한다. 
  - Main Routine으로 부터 각각의 Sub Routine으로 전송되어야 하고 
  - Sub Routine으로 부터 Main Routine으로 전송되어야 한다. 

### Channels

- Goroutine 사이에 대이터를 전송한다. 
- Channels은 타입이다. 
- 채널을 만들기 위해서 make()함수를 사용하라 

```go
package main

func main()  {
	c := make(chan int)
}
```

- Send and Receive using Arrow Operations 

- Send Data on Channel 

```go
package main

func main()  {
    c := make(chan int)
    
	c <- 3
}
```

- Receive Data from a channel 

```go
package main

func main()  {
    c := make(chan int)
  
	x := <- c
}
```

### Channel Example

```go
package main

import "fmt"

func prod(v1 int, v2 int, c chan int) {
  c <- v1 * v2
}

func main() {
  c := make(chan int)
  go prod(1, 2, c)
  go prod(3, 4, c)
  a := <-c
  b := <-c
  fmt.Println(a * b)
}
```

### Unbuffered Channel 

- 버퍼로 구성되어 있지 않는 채널은 데이터를 유지할 수 없음 
  - Sending(전송)은 데이터를 받을 때까지 차단한다. ( Block until data is received)
  - Receiving(수신)은 데이터가 전송되기 전까지 차단한다. ( Block until data is sent )

### Blocking And Synchronization 

- Channel Communication 는 동기화가 기본이다. 
- Blocking 은 커뮤니케이션을 위해서 대기상태가 되는 것과 같다, 

### Channel Capacity 

- 채널은 제한된 수의 오브젝트를 가질 수 있다. 
- Capacity는 오브젝트의 수리르 이야기 한다. 이것은 Transit 상태에서 유지할 수 있는 수를 의미한다.

```go
package main

import "fmt"

func main() {
  c := make(chan int, 3)
  fmt.Println(c)
}
```

Blocking은 기본적으로 좋은 것은 아니기 때문에 Capacity에 대해서 잘 고민해야한다. 왜나면 Block는 buffer가 풀일 경우에 발생하기 때문이다. 
수신 받는 것도 buffer가 없을 때 Block이 발생한다. 

### Channel Blocking, Receive 

![Group Routing Sync Wait](https://github.com/keepinmindsh/golang-books/blob/main/docs/assets/Concurrency_005.png)

위의 코드는 반드시 에러가 발생한다.   

우선 해당 코드가 Main 에서 실행될 때, T1, T2가 Go Routine에 의해서 동시에 실행될 때 T2는 버퍼가 비어있기 때문에 대기한다. 
그러다가 T1이 Channel에 버퍼를 넣어주면 T2는 그제서야 Buffer에서 채널의 값을 가져와서 설정한다.   
여기에서 문제가 생기는데 b := <- c 는 버퍼에서 공백인데 넣어주는 곳이 아무곳도 없기 때문에 에러가 발생하게 된다.   


반대의 경우에도 동일한 이슈가 생길 수 있는데 버버로 2개의 값을 보냈으나 받는 것이 오직 하나라면 이 경우에도 에러가 발생한다. 

### Use of Buffering 

- Sender와 Receiver가 정확히 동일한 속도로 처리될 필요가 없을 때 사용한다.
  - Buffer는 Producer와 Consumer 의 구조에서 도움이 된다.
    - Producer : Generating Data
    - Consumer : Use Data 
  - [Producer / Consumer 패턴](http://www.gisdeveloper.co.kr/?p=10325)

