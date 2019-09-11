package model

import (
	"chat/config"
	"chat/event"
	"fmt"
	"net"
	"runtime"
	"strings"
)

// 客户端抽象
type Client struct {
	addr string
	Name string
	ch   chan string
}

// 构造函数
func NewClient(addr string, name string, ch chan string) *Client {
	c := &Client{
		addr: addr,
		Name: name,
		ch:   ch,
	}
	c.online()
	return c
}

func (c *Client) online() {
	event.Trigger(event.ClientOnline, *c)
}

func (c *Client) offline() {
	event.Trigger(event.ClientOffline, *c)
}

// 处理请求
func (c *Client) handRequest(conn net.Conn) {
	buf := make([]byte, 2048)
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if nil != err {
			fmt.Println("请求处理失败:", err)
			return
		}
		received := string(buf[:n])
		received = strings.Trim(received, "\n")
		if len(received) == 5 && received[:4] == config.EXIT {
			c.offline()
			runtime.Goexit()
		} else {
			msg := fmt.Sprintf("%s say:\n %s", c.Name, received)
			fmt.Println(msg)
			ch <- msg
		}
	}
}

// 收到消息
func (c *Client) display(msg string) error {
	c.ch <- msg
	return nil
}
