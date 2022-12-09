package chapter9

import (
	"container/list"
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Header *Node
}

func LinkedList() {
	lst := list.New()
	lst.PushBack(100)
	lst.PushBack(101)
	lst.PushBack(102)

	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func (l *List) Constructor() *List {
	arr := []int{100, 101, 102}
	for _, v := range arr {
		l.Insert(v)
	}

	return &List{}
}

func (l *List) Insert(i int) {

	if l.Header == nil {

		l.Header = &Node{i, nil}
	} else {
		p := l.Header
		for p.Next != nil {
			p = p.Next
		}
		p.Next = &Node{i, nil}
	}
}

func (l *List) Show() {
	sl := l.Header
	for sl != nil {
		fmt.Println(sl.Val)
		sl = sl.Next
	}
}
