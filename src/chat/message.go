package chat

import (
	"fmt"
	"time"
)

// 消息结构体
type Message struct {
	Msg string
	To  *ClientsMap
	*Client
}

// 消息格式化字符串
func (m *Message) String() string {
	return fmt.Sprintf("[%s][%s] say: %s\n", time.Now().Format("2006-01-02 15:04:05"), m.Name, m.Msg)
}

// 消息格式化字节切片
func (m *Message) ByteSlice() []byte {
	return []byte(m.String())
}

// 消息构造
func NewMessage(cli *Client, msg string) *Message {
	return &Message{
		Msg:    msg,
		Client: cli,
	}
}
