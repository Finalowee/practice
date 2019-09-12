package day08

import (
	"fmt"
	"net"
	"net/http"
)

func TestHttpTcp() {
	listener, err := net.Listen("tcp", ":10086")
	if nil != err {
		fmt.Print(err)
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if nil != err {
		fmt.Print(err)
	}
	defer conn.Close()
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if nil != err {
		fmt.Print(err)
	}
	if 0 == n {
		return
	}
	fmt.Println(string(buf))
}

func handle(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
}

func TestHttp() {
	http.HandleFunc("/go", handle)
	http.ListenAndServe("127.0.0.1:10086", nil)
}
