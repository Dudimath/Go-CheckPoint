package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	current := l.Head

	if current == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}

func ListSize(l *List) int {
	var size int
	current := l.Head

	for current != nil {
		size++
		current = current.Next
	}
	return size
}
