package chat

import (
	"chat/event"
)

func init() {
	// 绑定服务端启动监听
	event.Register(event.ServerCreate, func(i interface{}) bool {
		//fmt.Printf("%+v\n", i)
		return true
	})
	// 用户上线
	event.Register(event.ClientOnline, func(i interface{}) bool {
		//cli := i.(model.Client)
		//onlineMsg := fmt.Sprintf("welcome %s", cli.Name)
		//err := cli.s.Broadcast(onlineMsg)
		//if nil != err {
		//	fmt.Println(err)
		//}
		return true
	})
	// 用户下线
	event.Register(event.ClientOffline, func(i interface{}) bool {
		//cli := i.(model.Client)
		//offlineMsg := fmt.Sprintf("%s offline", cli.Name)
		//err := cli.s.Broadcast(offlineMsg)
		//if nil != err {
		//	fmt.Println(err)
		//}
		return true
	})
}
