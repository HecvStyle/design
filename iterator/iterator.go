package iterator

type Iterator interface {
	Next() Member  // 下一个成员
	HasMore() bool // 是否有下一个
}

type memberIterator struct {
	class *Class
	index int
}

func (m *memberIterator) Next() Member {
	if m.index == -1 {
		m.index++
		return m.class.teacher
	}
	student := m.class.students[m.index]
	m.index++
	return student
}

func (m *memberIterator) HasMore() bool {
	return m.index < len(m.class.students)
}

type Iterable interface {
	CreateIterator() Iterator
}

type Class struct {
	name     string
	teacher  *Teacher
	students []*Student
}

func NewClass(name string, teacherName, teacherSubject string) *Class {
	return &Class{name: name, teacher: &Teacher{teacherName, teacherSubject}}
}

func (c *Class) CreateIterator() Iterator {
	return &memberIterator{class: c, index: -1}
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) AddStudent(students ...*Student) {
	c.students = append(c.students, students...)
}
