// 聊天服务器
package chat

import (
	"chat/config"
	"chat/event"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// 服务器抽象
type Server struct {
	ip            string
	port          int
	protocol      string
	onlineClients map[string]Client
	listener      net.Listener
	cliCnt        int
	cliMaxCnt     int
}

// 客户端抽象
type Client struct {
	addr string
	name string
	conn net.Conn
	s    *Server
}

// 初始化
func init() {
	// 绑定服务端启动监听
	event.Register(event.ServerCreate, func(i interface{}) bool {
		//fmt.Printf("%+v\n", i)
		return true
	})
	// 用户上线
	event.Register(event.ClientOnline, func(i interface{}) bool {
		cli := i.(Client)
		onlineMsg := fmt.Sprintf("welcome %s", cli.name)
		err := cli.s.Broadcast(onlineMsg)
		if nil != err {
			fmt.Println(err)
		}
		return true
	})
	// 用户下线
	event.Register(event.ClientOffline, func(i interface{}) bool {
		cli := i.(Client)
		offlineMsg := fmt.Sprintf("%s offline", cli.name)
		err := cli.s.Broadcast(offlineMsg)
		if nil != err {
			fmt.Println(err)
		}
		return true
	})
}

// 创建服务器实例
func NewServer() (*Server, error) {
	s := &Server{
		ip:            config.ServerIp,
		port:          config.ServerPort,
		protocol:      config.ServerProtocol,
		onlineClients: map[string]Client{},
	}
	return s, nil
}

// 开启监听
func (s *Server) listen() (err error) {
	if nil != s.listener {
		return
	}
	addr := s.ip + ":" + strconv.Itoa(s.port)
	listener, err := net.Listen(s.protocol, addr)
	if nil != err {
		fmt.Println("开启监听失败", addr, err)
		return err
	}
	s.listener = listener
	event.Trigger(event.ServerStartListen, s)
	return
}

// 开始处理请求
func (s *Server) Run(cliCnt int) {
	s.listen()
	s.cliMaxCnt = cliCnt
	for {
		if s.cliCnt > s.cliMaxCnt {
			time.Sleep(time.Second)
		} else {
			conn, err := s.listener.Accept()
			if nil != err {
				fmt.Println(err)
				return
			}
			// 初始化客户端
			cli := Client{
				addr: conn.RemoteAddr().String(),
				name: conn.RemoteAddr().String(),
				conn: conn,
				s:    s,
			}
			cli.online()
			go HandRequest(&cli)
		}
	}
}

// 处理请求
func HandRequest(cli *Client) {
	buf := make([]byte, 2048)
	defer cli.conn.Close()
	for {
		n, err := cli.conn.Read(buf)
		if nil != err {
			fmt.Println("请求处理失败:", err)
			return
		}
		received := string(buf[:n])
		received = strings.Trim(received, "\n")
		if len(received) == 5 && received[:4] == "EXIT" {
			cli.s.Broadcast(received)
			runtime.Goexit()
			return
			//fmt.Println("client:", remoteAddr, "offline")
		} else {
			//fmt.Println("replay:", replay)
			msg := fmt.Sprintf("%s say:\n %s", cli.name, string(received))
			cli.s.Broadcast(msg)
		}
	}
}

// 广播
func (s *Server) Broadcast(msg string) (err error) {
	for _, cli := range s.onlineClients {
		_, err := cli.conn.Write([]byte(msg + "\n"))
		if nil != err {
			fmt.Println(nil)
		}
	}
	return
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
