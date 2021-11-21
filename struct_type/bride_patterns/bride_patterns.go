package bride_patterns

import "fmt"

// 将抽象和实现分离开来
// 一个类存在两个以及两个以上的变化维度，可以通过组合的方式，将这些维度进行组合
// 现在有个需求，要做一个服务监控告警程序，告警方式可以是微信消息、短信消息、自动语言电话告警。同时告警级别有分，普通、重要、紧急。不同的告警级别对应不同的通知方式

type Message interface {
	SendMessage(message string)
}

type Notification interface {
	Send(message string)
}

//WeiXinNotification 微信
type WeiXinNotification struct{}

func NewWeiXinNotification() Notification {
	return &WeiXinNotification{}
}

func (*WeiXinNotification) Send(message string) {
	fmt.Printf("%s send [weixin] message \n", message)
}

//SMSNotification 短信
type SMSNotification struct{}

func NewSMSNotification() Notification {
	return &SMSNotification{}
}

func (*SMSNotification) Send(message string) {
	fmt.Printf("%s send [sms] message \n", message)
}

//PhoneNotification 手机
type PhoneNotification struct{}

func NewPhoneNotification() Notification {
	return &PhoneNotification{}
}

func (*PhoneNotification) Send(message string) {
	fmt.Printf("%s send [phone] message \n", message)
}

//CommonMessage 普通消息发送
type CommonMessage struct {
	notification Notification
}

func NewCommonMessage(notification Notification) *CommonMessage {
	return &CommonMessage{
		notification: notification,
	}
}

func (m *CommonMessage) SendMessage(message string) {
	m.notification.Send("[Common] notice " + message)
}

//ImportantMessage 重要消息发送
type ImportantMessage struct {
	notification Notification
}

func NewImportantMessage(notification Notification) *ImportantMessage {
	return &ImportantMessage{
		notification: notification,
	}
}

func (m *ImportantMessage) SendMessage(message string) {
	m.notification.Send("[Important] notice " + message)
}

//UrgencyMessage 自动化消息发送
type UrgencyMessage struct {
	notification Notification
}

func NewUrgencyMessage(notification Notification) *UrgencyMessage {
	return &UrgencyMessage{
		notification: notification,
	}
}

func (m *UrgencyMessage) SendMessage(message string) {
	m.notification.Send("[Urgency] notice " + message)
}
