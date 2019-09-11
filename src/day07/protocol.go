package day07

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
)

func NewTCPServer(addr string) {

	listener, err := net.Listen("tcp", addr)
	if nil != err {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if nil != err {
			fmt.Println(err)
			return
		}
		go handleRequest(conn)
	}
	return
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	remoteAddr := conn.RemoteAddr().String()
	conn.Write([]byte("Hello, " + remoteAddr))
	for {
		//fmt.Println("client:", remoteAddr, "online")
		n, err := conn.Read(buf)
		if nil != err {
			fmt.Println(err)
			return
		}
		received := string(buf[:n])
		fmt.Println("[log]Server received:", received, ",len:", len(received))
		replay := strings.ToUpper(string(buf[:n]))
		if len(replay) == 5 && replay[:4] == "EXIT" {
			conn.Write([]byte("bye"))
			runtime.Goexit()
			return
			//fmt.Println("client:", remoteAddr, "offline")
		} else {
			//fmt.Println("replay:", replay)
			conn.Write([]byte(replay))
		}
	}
}

func NewTcpClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if nil != err {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if nil != err {
			fmt.Println(err)
			return
		}
		received := string(buf[:n])
		fmt.Println("received:", received, ",len:", len(received))
		n, err = os.Stdin.Read(buf)
		if nil != err {
			fmt.Println(err)
			return
		}
		conn.Write(buf[:n])
	}
}
