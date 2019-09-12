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
	chBroadcast   chan Message
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
		onlineClients: make(map[string]Client),
		chBroadcast:   make(chan Message),
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
	defer func() {
		err = s.listener.Close()
	}()
	s.cliMaxCnt = cliCnt
	go s.processMsgBroadcast()
	s.processRequest() // 阻塞
	return
}

func (s *Server) bindListener() {
	event.Register(event.ClientOnline, s.online)
	event.Register(event.ClientOnMessage, s.send)
	//event.Register(event.ClientOffline, s.offline)
}

func (s *Server) processMsgBroadcast() {
	for {
		msg := <-s.chBroadcast
		fmt.Println(msg)
		if nil == msg.to {
			msg.to = s.onlineClients
		}
		for _, cli := range msg.to {
			cli.ch <- msg
		}
	}
}

func (s *Server) processRequest() {
	for n := 1; true; n++ {
		if s.cliCnt > s.cliMaxCnt {
			time.Sleep(time.Second)
			continue
		}
		conn, err := s.listener.Accept()
		if nil != err {
			fmt.Println(err)
			continue
		}
		// 初始化客户端
		addr := conn.RemoteAddr().String()
		cli := NewClient(addr, "user_"+strconv.Itoa(n))
		go cli.handMessage(conn)
	}
}

func (s *Server) online(i interface{}) bool {
	s.cliCnt++
	cli := i.(Client)
	s.onlineClients[cli.addr] = cli
	fmt.Printf("user online:\n %+v\n", cli)
	cli.SendMessage("Hello, everyone!", nil)
	return true
}

func (s *Server) send(i interface{}) bool {
	onlineMsg := i.(*Message)
	//fmt.Println(onlineMsg)
	s.chBroadcast <- *onlineMsg
	return true
}

//func (s *Server) offline(i interface{}) bool {
//	s.cliCnt--
//	cli := i.(Client)
//
//	offlineMsg := fmt.Sprintf("%s offline", cli.Name)
//	_ = cli.s.Broadcast(offlineMsg)
//	delete(s.onlineClients, cli.addr)
//	return true
//}
