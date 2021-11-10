package bst

import (
	"fmt"
	"learn/queue"
	"learn/stack"
)

//binary search tree

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

//Remove 以指针的模式移除
func (b *BST) Remove(v int) bool {
	return b.remove(&b.root, v)
}

//Remove1 以父节点的模式移除
func (b *BST) Remove1(v int) bool {
	return b.remove1(nil, b.root, v)
}

//引入父节点模式
func (b *BST) remove1(front, root *Node, v int) bool {
	setNode := func(front, root, node *Node) *Node {
		if front.left.value == root.value {
			front.left = node
			return front.left
		} else {
			front.right = node
			return front.right
		}
	}
	if root == nil {
		return false
	} else if v < root.value {
		return b.remove1(root, root.left, v)
	} else if v > root.value {
		return b.remove1(root, root.right, v)
	}
	//find node && child empty
	if root.left == nil && root.right == nil {
		setNode(front, root, nil)
		return true
	}
	if root.left == nil && root.right != nil {
		setNode(front, root, root.right)
		return true
	}
	if root.right == nil && root.left != nil {
		setNode(front, root, root.left)
		return true
	}
	// root have full child
	minNode := b.findMinNode(root.right)
	return b.remove1(setNode(front, root, &Node{
		value: minNode.value,
		left:  root.left,
		right: root.right,
	}), root.right, minNode.value)
}

//指针的指针模式
func (b *BST) remove(root **Node, v int) bool {
	if root == nil {
		return false
	} else if v < (*root).value {
		return b.remove(&(*root).left, v)
	} else if v > (*root).value {
		return b.remove(&(*root).right, v)
	}
	//find node && child empty
	if (*root).left == nil && (*root).right == nil {
		*root = nil
		return true
	}
	if (*root).left == nil && (*root).right != nil {
		*root = (*root).right
		return true
	}
	if (*root).right == nil && (*root).left != nil {
		*root = (*root).left
		return true
	}
	// root have full child
	minNode := b.findMinNode((*root).right)
	(*root).value = minNode.value
	return b.remove(&(*root).right, minNode.value)
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
		return b.search(root.left, v)
	} else if v > root.value {
		return b.search(root.right, v)
	}
	return true
}

func (b *BST) Min() (int, bool) {
	return b.min(b.root)
}

func (b *BST) min(root *Node) (int, bool) {
	if root.left == nil {
		return root.value, true
	}
	return b.min(root.left)
}

func (b *BST) Max() (int, bool) {
	return b.max(b.root)
}

func (b *BST) max(root *Node) (int, bool) {
	if root.right == nil {
		return root.value, true
	}
	return b.max(root.right)
}

//前序遍历栈实现
func (b *BST) PreOrderTraveralByStack() {
	node, stack := b.root, stack.NewStack(20)
	for node != nil || stack.Len() != 0 {
		//迭代访问左孩子,入栈
		for node != nil {
			fmt.Printf(" %d ", node.value)
			stack.Push(node)
			node = node.left
		}
		if stack.Len() != 0 {
			n, _ := stack.Pop()
			node = n.(*Node)
			node = node.right
		}
	}
	fmt.Println()
}

//前序遍历递归实现
func (b *BST) PreOrderTraveral() {
	b.preOrderTraveral(b.root)
	fmt.Printf("\n")
}

func (b *BST) preOrderTraveral(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf(" %d ", root.value)
	b.preOrderTraveral(root.left)
	b.preOrderTraveral(root.right)
}

//中序遍历 递归
func (b *BST) InOrderTraveral() {
	b.inOrderTraveral(b.root)
	fmt.Printf("\n")
}

func (b *BST) inOrderTraveral(root *Node) {
	if root == nil {
		return
	}
	b.preOrderTraveral(root.left)
	fmt.Printf(" %d ", root.value)
	b.preOrderTraveral(root.right)
}

func (b *BST) RearOrderTraveral() {
	b.rearOrderTraveral(b.root)
	fmt.Printf("\n")
}

func (b *BST) rearOrderTraveral(root *Node) {
	if root == nil {
		return
	}
	b.preOrderTraveral(root.left)
	b.preOrderTraveral(root.right)
	fmt.Printf(" %d ", root.value)
}

func (b *BST) Print() {
	out, qe := make([]string, 0), queue.NewQueue(100)
	qe.EnQueue(b.root)
	for qe.Len() != 0 {
		n, str := qe.Len(), ""
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
	maxLen, padding := len(out[len(out)-1]), ""
	for _, s := range out {
		for i := 0; i < (maxLen-len(s))/2; i++ {
			padding += " "
		}
		fmt.Printf("%s \n\n", padding+s)
		padding = ""
	}
}
