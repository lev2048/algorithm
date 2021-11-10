package bst

import (
	"fmt"
	"testing"
)

func TestBSTRemove1(t *testing.T) {
	bt := new(BST)
	data := [6]int{5, 3, 6, 2, 4, 7}
	for _, v := range data {
		bt.Insert(v)
	}
	fmt.Println("Data: ", data)
	fmt.Printf("Print BST: \n\n")
	bt.Print()

	fmt.Printf("Remove Item 3 And Print: \n\n")
	bt.Remove(3)
	bt.Print()
}

func TestBSTRemove2(t *testing.T) {
	bt := new(BST)
	data := [9]int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range data {
		bt.Insert(v)
	}
	fmt.Println("Data: ", data)
	fmt.Printf("Print BST: \n\n")
	bt.Print()

	fmt.Printf("Remove Item 3 And Print: \n\n")
	bt.Remove(3)
	bt.Print()
}

func TestBSTGetMin(t *testing.T) {
	bt := new(BST)
	data := [9]int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range data {
		bt.Insert(v)
	}
	if v, _ := bt.Min(); v != 1 {
		t.Fatal("fail: ", v)
	}
}

func TestBSTGetMax(t *testing.T) {
	bt := new(BST)
	data := [9]int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range data {
		bt.Insert(v)
	}
	if v, _ := bt.Max(); v != 14 {
		t.Fatal("fail: ", v)
	}
}

func TestBSTTraveral(t *testing.T) {
	bt := new(BST)
	data := [9]int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range data {
		bt.Insert(v)
	}
	fmt.Println("前序遍历 递归")
	bt.PreOrderTraveral()
	fmt.Println()

	fmt.Println("前序遍历 Stack")
	bt.PreOrderTraveralByStack()
	fmt.Println()

	fmt.Println("中序遍历")
	bt.InOrderTraveral()
	fmt.Println()

	fmt.Println("后序遍历")
	bt.RearOrderTraveral()
	fmt.Println()
}
