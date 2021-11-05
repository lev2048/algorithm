package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type SingleList struct {
	size int
	head *Node
	tail *Node
}

func NewSingleList() *SingleList {
	return &SingleList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (l *SingleList) PushBack(v interface{}) *Node {
	node := &Node{
		Data: v,
		Next: nil,
	}
	if l.size == 0 {
		l.head, l.tail = node, node
	} else {
		if l.head.Next == nil {
			l.head.Next, l.tail = node, node
		} else {
			l.tail.Next, l.tail = node, node
		}
	}
	l.size++
	return node
}

func (l *SingleList) PushFront(v interface{}) *Node {
	node := &Node{
		Data: v,
		Next: l.head,
	}
	l.head = node
	l.size++
	return node
}

func (l *SingleList) InsertAfter(mark int, v interface{}) *Node {
	node := &Node{
		Data: v,
		Next: nil,
	}
	if mark == 0 {
		node.Next = l.head
		l.head = node
	} else if mark == l.size {
		l.tail.Next = node
	} else if mark > 0 && mark < l.size {
		target := l.head
		for i := 1; i < mark; i++ {
			target = target.Next
		}
		node.Next = target.Next
		target.Next = node
	} else {
		return nil
	}
	l.size++
	return node
}

func (l *SingleList) Update(mark int, data interface{}) {
	if tn := l.Get(mark); tn != nil {
		tn.Data = data
	}
}

func (l *SingleList) Remove(mark int) {
	if mark == 1 {
		l.head = l.head.Next
	} else {
		target := l.head
		for i := 1; i < mark-1; i++ {
			target = target.Next
		}
		target.Next = target.Next.Next
	}
	l.size--
}

func (l *SingleList) Get(mark int) *Node {
	if mark == 1 {
		return l.head
	} else if mark == l.size {
		return l.tail
	} else if mark > 0 && mark < l.size {
		target := l.head
		for i := 1; i < mark; i++ {
			target = target.Next
		}
		return target
	} else {
		return nil
	}
}

func (l *SingleList) Len() int {
	return l.size
}

func main() {
	sl := NewSingleList()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range data {
		sl.PushBack(data[len(data)-i-1])
	}
	outPutList := func(mark string) {
		fmt.Printf("%s :", mark)
		for i := 1; i < sl.Len(); i++ {
			fmt.Printf("%d ", sl.Get(i).Data)
		}
		fmt.Println()
	}
	outPutList("Print")
	sl.InsertAfter(2, 2)
	outPutList("InsterAfter")
	sl.PushFront(1)
	outPutList("PushFront")
	sl.Remove(1)
	sl.Remove(2)
	outPutList("Remove 1 2")
	sl.Update(2, 3)
	outPutList("Update 2=>3")
}
