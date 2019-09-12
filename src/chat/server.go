package chat

import (
	"fmt"
	"net"
	"strings"
)

// 广播通道
var chBroadcast MessageChannel

// 客户端表
var cliMap ClientsMap

// 主函数体
func main() {
	addr := ":10086"
	listener, err := net.Listen("tcp", addr)
	if nil != err {
		fmt.Println("net.listen():", err)
	}
	defer listener.Close()
	chBroadcast = make(chan Message)
	cliMap = make(ClientsMap)
	// 发送消息到客户端
	go ProcessMessage()
	// 接受请求
	for n := 1; true; n++ {
		conn, err := listener.Accept()
		if nil != err {
			fmt.Println("listener.Accept():", err)
		}
		cli := NewClient(conn, "")
		go ProcessRequest(cli, conn)
	}
}

func ProcessMessage() {
	for msg := range chBroadcast {
		if msg.To == nil {
			msg.To = &cliMap
		}
		for _, cli := range *msg.To {
			cli.ChanMsg <- msg
		}
	}
}

func ProcessRequest(cli *Client, conn net.Conn) {
	// 向客户端发送消息
	go func() {
		for {
			msg := <-cli.ChanMsg
			conn.Write(msg.ByteSlice())
		}
	}()
	// 处理客户端发来的消息
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if nil != err {
				fmt.Println(err)
			}
			if 0 == n {
				fmt.Println("user offline:", cli.Name)
				delete(cliMap, cli.Addr)
				chBroadcast <- *NewMessage(cli, "bye")
				return
			}
			chBroadcast <- *ParseInput(cli, buf[:n-1])
		}
	}()
	// 处理客户端发来的命令
	for {
		select {
		case msg := <-cli.ChanCmd:

		}
	}
}

func ParseInput(cli *Client, b []byte) *Message {
	str := string(b)
	m := NewMessage(cli, str)
	if str[0] == '/' {
		m.Msg = parseCmd(cli, str[1:])
		m.To = &ClientsMap{cli.Addr: cli}
	}
	return m
}

func parseCmd(cli *Client, str string) string {
	var ret string
	args := strings.Split(str, " ")
	switch args[0] {
	case "who":
		ret = "users:\n"
		for _, cli := range cliMap {
			ret += fmt.Sprintf("addr: %s; name: %s \n", cli.Addr, cli.Name)
		}
	case "rename":
		if len(args) > 1 {
			cli.Name = args[1]
			ret = "success"
		}
	case "exit":
		cli.ChanCmd <- *NewMessage(cli, "EXIT")
	}
	return ret
}
