package protocol_server

import (
	f "fmt"
	"net"
	"net/rpc"
)

type Calc int // RPC 서버에 등록하기 위해 임의의 타입으로 정의

type Args struct { // 매개변수
	A, B int
}

type Reply struct { // 리턴값
	C int
}

func Client() {

	f.Println("rpc server running at 6000 port")

	rpc.Register(new(Calc)) // Calc 타입의 인스턴스를 서버의 등록

	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		defer conn.Close()

		go rpc.ServeConn(conn)
	}

}

func (c *Calc) Sum(args Args, reply *Reply) error {

	reply.C = args.A + args.B // 두 값을 더하여 리턴값 구조체에 넣어줌
	return nil

}
