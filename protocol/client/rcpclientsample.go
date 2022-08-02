package protocol_client

import (
	f "fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func Server() {

	client, err := rpc.Dial("tcp", "127.0.0.1:6000") // RPC 서버 접ㅈ속
	if err != nil {
		f.Println(err)
		return
	}
	defer client.Close()

	// 동기 호출
	args := &Args{1, 2}
	reply := new(Reply)

	err = client.Call("Calc.Sum", args, reply) // 함수 호출
	if err != nil {
		f.Println(err)
		return
	}
	f.Println(reply.C)

	// 비동기 호출
	args.A = 4
	args.B = 9

	sumCall := client.Go("Calc.Sum", args, reply, nil) // 고루틴으로 호출
	<-sumCall.Done
	f.Println(reply.C)

}
