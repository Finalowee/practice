package chat

import "net"

// 客户端结构体
type Client struct {
	Addr    string
	Name    string
	ChanMsg chan Message
	ChanCmd chan Message
}

// 客户端构造
func NewClient(conn net.Conn, name string) *Client {
	addr := conn.RemoteAddr().String()
	if "" == name {
		name = addr
	}
	cli := &Client{
		addr,
		name,
		make(MessageChannel),
		make(MessageChannel),
	}
	cliMap[addr] = cli
	msg := NewMessage(cli, "Hello everyone!")
	chBroadcast <- *msg
	return cli
}
