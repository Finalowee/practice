package chat

import (
	"chat/event"
	"fmt"
	"net"
	"runtime"
	"strings"
)

// 客户端抽象
type Client struct {
	addr string
	name string
	conn net.Conn
	s    *Server
}

func (c *Client) online() {
	c.s.cliCnt++
	c.s.onlineClients[c.addr] = *c
	event.Trigger(event.ClientOnline, *c)
}

func (c *Client) offline() {
	c.s.cliCnt--
	delete(c.s.onlineClients, c.addr)
	event.Trigger(event.ClientOffline, *c)
}

// 处理请求
func (c *Client) HandRequest() {
	buf := make([]byte, 2048)
	defer c.conn.Close()
	for {
		n, err := c.conn.Read(buf)
		if nil != err {
			fmt.Println("请求处理失败:", err)
			return
		}
		received := string(buf[:n])
		received = strings.Trim(received, "\n")
		if len(received) == 5 && received[:4] == "EXIT" {
			_ = c.s.Broadcast(received)
			runtime.Goexit()
			return
			//fmt.Println("client:", remoteAddr, "offline")
		} else {
			//fmt.Println("replay:", replay)
			msg := fmt.Sprintf("%s say:\n %s", c.name, string(received))
			_ = c.s.Broadcast(msg)
		}
	}
}
