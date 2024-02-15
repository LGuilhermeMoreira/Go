package list

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type linkedList struct {
	Head   *Node
	Lenght int
}

func Start() *linkedList {
	return &linkedList{
		Head:   nil,
		Lenght: 0,
	}
}

func (l *linkedList) Add(value interface{}) {
	l.Lenght = +1

	node := Node{Value: value, Next: nil}

	if l.Head == nil {
		l.Head = &Node{Value: "Head", Next: &node}
		return
	}

	head := l.Head

	for head.Next != nil {
		head = head.Next
	}

	head.Next = &node
}

func (l *linkedList) Show() {
	head := l.Head

	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}
