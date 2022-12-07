package observer

import "fmt"

type Subscriber interface {
	Name() string
	Update(message string)
}

type shortMessage struct {
}

func (s *shortMessage) Name() string {
	return "手机短信"
}

func (s *shortMessage) Update(message string) {
	fmt.Printf("通过%s发送消息：%s\n", s.Name(), message)
}

type email struct {
}

func (e *email) Name() string {
	return "电子邮件"
}

func (e *email) Update(message string) {
	fmt.Printf("通过%s发送消息：%s\n", e.Name(), message)
}

type telephone struct {
}

func (t *telephone) Name() string {
	return "电话"
}

func (t *telephone) Update(message string) {
	fmt.Printf("通过%s告知消息：%s\n", t.Name(), message)
}
