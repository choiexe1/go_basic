package pointerprac

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil}
}

func (l *LinkedList) Insert(value int) {
	newNode := NewNode(value)

	if l.Head == nil {
		l.Head = newNode
	} else {
		prev := l.Head
		for prev.Next != nil {
			prev = prev.Next
		}

		prev.Next = newNode
	}
}

func (l *LinkedList) Delete(value int) bool {
	if l.Head == nil {
		return false
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return true
	}

	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.Value == value {
			prev.Next = curr.Next
			return true
		}

		prev = curr
		curr = curr.Next
	}

	return false
}

func (l *LinkedList) Find(value int) bool {
	if l.Head == nil {
		return false
	}

	curr := l.Head

	for curr != nil {
		if curr.Value == value {
			return true
		}
		curr = curr.Next
	}

	return false
}

func (l *LinkedList) String() string {
	curr := l.Head
	var result strings.Builder

	for curr != nil {
		fmt.Fprintf(&result, "%d -> ", curr.Value)
		curr = curr.Next
	}

	fmt.Fprintf(&result, "nil")

	return result.String()
}
