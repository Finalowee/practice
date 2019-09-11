// 聊天服务器
package model

import (
	"chat/config"
	"chat/event"
	"fmt"
	"net"
	"strconv"
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

// 初始化
func init() {
}

// 创建服务器实例
func NewServer() (*Server, error) {
	s := &Server{
		ip:            config.ServerIp,
		port:          config.ServerPort,
		protocol:      config.ServerProtocol,
		onlineClients: map[string]Client{},
	}
	s.bindListener()
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
func (s *Server) Run(cliCnt int) (err error) {
	err = s.listen()
	if nil != err {
		return
	}
	s.cliMaxCnt = cliCnt
	var ch chan string
	n := 0
	for {
		n++
		if s.cliCnt > s.cliMaxCnt {
			time.Sleep(time.Second)
		} else {
			conn, err := s.listener.Accept()
			if nil != err {
				return err
			}
			// 初始化客户端
			addr := conn.RemoteAddr().String()
			cli := NewClient(addr, "user_"+strconv.Itoa(n), make(chan string))
			go cli.handRequest(conn)
		}
	}
}

// 广播
func (s *Server) Broadcast(msg string) (err error) {
	for _, cli := range s.onlineClients {
		err = cli.display(msg)
	}
	return
}

func (s *Server) online(i interface{}) bool {
	s.cliCnt++
	cli := i.(Client)
	s.onlineClients[cli.addr] = cli

	onlineMsg := fmt.Sprintf("welcome %s", cli.Name)
	_ = s.Broadcast(onlineMsg)
	return true
}
func (s *Server) offline(i interface{}) bool {
	s.cliCnt--
	cli := i.(Client)

	offlineMsg := fmt.Sprintf("%s offline", cli.Name)
	_ = cli.s.Broadcast(offlineMsg)
	delete(s.onlineClients, cli.addr)
	return true
}
func (s *Server) bindListener() {
	event.Register(event.ClientOnline, s.online)
	event.Register(event.ClientOffline, s.offline)
}
