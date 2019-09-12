// 自定义类型
package chat

// 客户端字典，用户保存客户端集合
type ClientsMap map[string]*Client

// 通道
type MessageChannel chan Message
