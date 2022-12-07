package iterator

import "fmt"

type Member interface {
	// 成员的描述
	Des() string
}

type Teacher struct {
	name    string
	subject string
}

func (t *Teacher) Des() string {
	return fmt.Sprintf("%s 老师负责教授 %s 科目", t.name, t.subject)
}

func NewTeacher(name string, subject string) *Teacher {
	return &Teacher{name: name, subject: subject}
}

type Student struct {
	name     string
	sumScore int
}

func NewStudent(name string, sumScore int) *Student {
	return &Student{name: name, sumScore: sumScore}
}

func (s *Student) Des() string {
	return fmt.Sprintf("学生名字:%s，总分:%d", s.name, s.sumScore)
}
