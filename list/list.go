package list

type Node struct {
	Data interface{}
	Prev *Node
	Next *Node
}

type List struct {
	size int
	head *Node
	tail *Node
}

func (l *List) Len() int { return l.size }

func (l *List) Get(mark int) *Node {
	if mark == 1 {
		return l.head
	} else if mark == l.size {
		return l.tail
	} else if mark > 1 && mark < l.size {
		target := l.head
		for i := 1; i < mark; i++ {
			target = target.Next
		}
		return target
	} else {
		return nil
	}
}

func (l *List) PushFront(v interface{}) *Node {
	node := &Node{
		Data: v,
		Prev: nil,
		Next: nil,
	}
	if l.size == 0 {
		l.head, l.tail = node, node
	} else {
		l.head.Prev, node.Next = node, l.head
		l.head = node
	}
	l.size++
	return node
}

func (l *List) PushBack(v interface{}) *Node {
	node := &Node{
		Data: v,
		Prev: nil,
		Next: nil,
	}
	if l.size == 0 {
		l.head, l.tail = node, node
	} else {
		l.tail.Next, node.Prev = node, l.tail
		l.tail = node
	}
	l.size++
	return node
}

func (l *List) Insert(mark int, v interface{}) (node *Node) {
	if mark == l.size {
		node = l.PushBack(v)
	} else if mark >= 1 && mark < l.size {
		node = &Node{
			Data: v,
			Prev: nil,
			Next: nil,
		}
		target := l.head
		for i := 1; i < mark; i++ {
			target = target.Next
		}
		node.Prev, node.Next = target, target.Next
		target.Next, target.Next.Prev = node, node
	} else {
		return
	}
	l.size++
	return
}

func (l *List) Update(mark int, v interface{}) *Node {
	if node := l.Get(mark); node != nil {
		newNode := &Node{
			Data: v,
			Prev: node.Prev,
			Next: node.Next,
		}
		node.Prev.Next, node.Next.Prev = newNode, newNode
		return newNode
	}
	return nil
}

func (l *List) Remove(mark int) {
	if mark == 1 {
		l.head.Next.Prev = nil
		l.head = l.head.Next
	} else if mark == l.size {
		l.tail.Prev.Next, l.tail.Next = nil, nil
		l.tail = l.tail.Prev
	} else if mark > 1 && mark < l.size {
		target := l.head
		for i := 1; i < mark; i++ {
			target = target.Next
		}
		target.Prev.Next = target.Next
		target.Next.Prev = target.Prev
	}
	l.size--
}
