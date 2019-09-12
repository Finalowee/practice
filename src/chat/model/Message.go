package model

import (
	"fmt"
	"time"
)

type Message struct {
	body string
	from *Client
	tag  int
	to   map[string]Client
}

func NewMessage(msg string, from *Client, to map[string]Client) *Message {
	cli := &Message{
		body: msg,
		from: from,
		to:   to,
		tag:  0,
	}
	return cli
}

func (m *Message) String() (str string) {
	str = fmt.Sprintf("[%s][%s]say:\n%s\n", time.Now().Format("2006-01-02 15:04:05"), m.from.Name, m.body)
	return
}
