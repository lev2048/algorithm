package bst

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (b *BST) Insert(v int) {
	newNode := &Node{value: v}
	if b.root == nil {
		b.root = newNode
	}
	b.insert(b.root, newNode)
}

//todo 重复元素引入计数,当前忽略
func (b *BST) insert(root, newNode *Node) {
	if newNode.value < root.value {
		if root.left == nil {
			root.left = newNode
		} else {
			b.insert(root.left, newNode)
		}
	} else if newNode.value > root.value {
		if root.right == nil {
			root.right = newNode
		} else {
			b.insert(root.right, newNode)
		}
	}
	//重复
}

func (b *BST) Remove(v int) bool {
	return b.remove(b.root, v)
}

func (b *BST) remove(root *Node, v int) bool {
	if root == nil {
		return false
	} else if v < root.value {
		b.remove(root.left, v)
	} else if v > root.value {
		b.remove(root.right, v)
	}
	//find node && child empty
	if root.left == nil && root.right == nil {
		root = nil
		return true
	}
	if root.left == nil && root.right != nil {
		root = root.right
		return true
	}
	if root.right == nil && root.left != nil {
		root = root.left
		return true
	}
	// root have full child
	minNode := b.findMinNode(root.right)
	b.remove(root.right, minNode.value)
	root = minNode
	return true
}

func (b *BST) findMinNode(root *Node) *Node {
	if root.left == nil {
		return root
	} else {
		return b.findMinNode(root.left)
	}
}

func (b *BST) Search(v int) bool {
	return b.search(b.root, v)
}

func (b *BST) search(root *Node, v int) bool {
	if root == nil {
		return false
	} else if v < root.value {
		b.search(root.left, v)
	} else if v > root.value {
		b.search(root.right, v)
	}
	return true
}

func (b *BST) Print() {
	out := make([]string, 0)
	qe := NewQueue(100)
	qe.EnQueue(b.root)
	for qe.len != 0 {
		n := qe.len
		str := ""
		for i := 0; i < n; i++ {
			if node, err := qe.DeQueue(); err != nil {
				return
			} else {
				if node == nil {
					str += fmt.Sprintf(" %s ", "X")
					continue
				}
				if node.(*Node).left != nil {
					qe.EnQueue(node.(*Node).left)
				} else {
					qe.EnQueue(nil)
				}
				if node.(*Node).right != nil {
					qe.EnQueue(node.(*Node).right)
				} else {
					qe.EnQueue(nil)
				}
				str += fmt.Sprintf(" %d ", node.(*Node).value)
			}
		}
		out = append(out, str)
	}
	out = out[:len(out)-1]
	w := len(out[len(out)-1])
	for _, s := range out {
		pad := ""
		for i := 0; i < (w-len(s))/2; i++ {
			pad += " "
		}
		fmt.Printf("%s \n", pad+s)
	}
}

func main() {
	bst := new(BST)
	data := [9]int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range data {
		bst.Insert(v)
	}
	bst.Print()
	fmt.Println(bst.Search(5))
	bst.Remove(6)
	bst.Print()
}

type Queue struct {
	len   int
	size  int
	rear  int
	front int
	data  []interface{}
}

func NewQueue(size int) *Queue {
	return &Queue{
		front: 0,
		rear:  0,
		size:  size,
		data:  make([]interface{}, size),
	}
}

func (q *Queue) EnQueue(v interface{}) error {
	if (q.rear+1)%q.size == q.front {
		return errors.New("queue is full")
	}
	q.data[q.rear] = v
	q.rear = (q.rear + 1) % q.size
	q.len++
	return nil
}

func (q *Queue) DeQueue() (interface{}, error) {
	if q.rear == q.front {
		return nil, errors.New("queue is empty")
	}
	item := q.data[q.front]
	q.front, q.data[q.front] = (q.front+1)%q.size, nil
	q.len--
	return item, nil
}
