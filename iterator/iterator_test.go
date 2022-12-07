package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	class := NewClass("三年二班", "Mr Wang", "Chinese")
	class.AddStudent(NewStudent("Lily", 99), NewStudent("Tom", 80), NewStudent("jon", 88))
	fmt.Printf("%s 成员如下:\n", class.name)
	classIterator := class.CreateIterator()
	for classIterator.HasMore() {
		member := classIterator.Next()
		fmt.Println(member.Des())
	}
}
