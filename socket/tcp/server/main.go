package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type SHandler interface {
	handle(conn net.Conn)
	accept() (net.Conn, error)
}

type serverHandler struct {
	ln net.Listener
}

func (s *serverHandler) handle(conn net.Conn) {
	buf := make([]byte, 4096)

	for {
		count, err := conn.Read(buf)
		if count == 0 || err != nil {
			_ = conn.Close()
			break
		}

		//3.3 根据输入流进行逻辑处理
		//buf数据 -> 去两端空格的string
		inStr := strings.TrimSpace(string(buf[0:count]))
		//去除 string 内部空格
		cInputs := strings.Split(inStr, " ")
		//获取 客户端输入第一条命令
		fCommand := cInputs[0]

		fmt.Println("客户端传输->" + fCommand)

		switch fCommand {
		case "ping":
			conn.Write([]byte("服务器端回复-> pong\n"))
		case "hello":
			conn.Write([]byte("服务器端回复-> world\n"))
		default:
			conn.Write([]byte("服务器端回复" + fCommand + "\n"))
		}

	}
}

func (s *serverHandler) accept() (net.Conn, error) {
	conn, err := s.ln.Accept()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func newSHandler(l net.Listener) SHandler {
	return &serverHandler{l}
}

func main() {
	l, err := net.Listen("tcp", ":8777")
	if err != nil {
		log.Fatalf("listen tcp failed: %+v", err)
	}

	fmt.Printf("start tcp listen on %+v\n", 8777)

	sh := newSHandler(l)

	for {
		conn, err := sh.accept()
		if err != nil {
			fmt.Println("to do reconn")
		}

		go sh.handle(conn)
	}

}
