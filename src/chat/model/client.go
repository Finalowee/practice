package model

import (
	"chat/event"
	"net"
)

// 客户端抽象
type Client struct {
	addr string
	Name string
	ch   chan Message
	ic   chan Message
}

// 构造函数
func NewClient(addr string, name string) *Client {
	c := &Client{
		addr: addr,
		Name: name,
		ch:   make(chan Message),
	}
	return c
}

func (c *Client) online() {
	event.Trigger(event.ClientOnline, *c)
}

func (c *Client) offline() {
	event.Trigger(event.ClientOffline, *c)
}

// 处理请求
/*
请求处理方法

@param conn 连接句柄
@ic 	输入管道
*/
func (c *Client) handMessage(conn net.Conn) {
	defer conn.Close()
	go func() {
		msg := <-c.ch
		conn.Write([]byte(msg.String()))
	}()
	c.online()
	for true {

	}
	//buf := make([]byte, 2048)
	//for {
	//	n, err := conn.Read(buf)
	//	if nil != err {
	//		fmt.Println("请求处理失败:", err)
	//		return
	//	}
	//	received := string(buf[:n])
	//	received = strings.Trim(received, "\n")
	//	if len(received) == 5 && received[:4] == config.EXIT {
	//		c.offline()
	//		runtime.Goexit()
	//	} else {
	//		msg := NewMessage(received, c, nil)
	//		fmt.Println(msg)
	//		c.ch <- *msg
	//	}
	//}
}

//// 收到消息
//func (c *Client) display(msg string) error {
//	c.ch <- msg
//	return nil
//}

func (c *Client) SendMessage(body string, to map[string]Client) {
	event.Trigger(event.ClientOnMessage, NewMessage(body, c, to))
}
