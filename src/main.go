package main

import (
	"chat/model"
	"fmt"
)

func main() {
	//day05.Copy("./src/day05/copy.go", "./src/day05/copy-bak.go")
	//day06.TestGoSched()

	//go day07.NewTCPServer("127.0.0.1:8888")
	//day07.NewTcpClient("127.0.0.1:8888")
	server, err := model.NewServer()
	if nil != err {
		fmt.Println(err)
	}
	_ = server.Run(3)

}
