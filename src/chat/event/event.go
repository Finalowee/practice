// 事件处理组件
package event

type Id int

type Handle func(interface{}) (flag bool)

var events map[Id][]Handle

func init() {
	events = map[Id][]Handle{}
}

// 注册事件处理函数
func Register(id Id, callable Handle) {
	handles := events[id]
	handles = append(handles, callable)
	events[id] = handles
}

// 触发事件
func Trigger(id Id, dat interface{}) {
	go func() {
		handles := events[id]
		for _, handle := range handles {
			ret := handle(dat)
			// 返回 false 则事件停止冒泡
			if !ret {
				break
			}
		}
	}()
}
