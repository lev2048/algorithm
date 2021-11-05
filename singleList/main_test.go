package main

import (
	"fmt"
	"testing"
)

func TestSingleList(t *testing.T) {
	sl := NewSingleList()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	outPut := func() string {
		res := ""
		for i := 1; i <= sl.Len(); i++ {
			res += fmt.Sprintf("%d ", sl.Get(i).Data)
		}
		return res
	}

	t.Run("pushBack", func(t *testing.T) {
		for _, v := range data {
			sl.PushBack(v)
		}
		if outPut() != "1 2 3 4 5 6 7 8 9 " {
			t.Fatal("fail", outPut())
		}
	})

	t.Run("PushFront", func(t *testing.T) {
		sl.PushFront(5)
		if outPut() != "5 1 2 3 4 5 6 7 8 9 " {
			t.Fatal("fail", outPut())
		}
	})

	t.Run("InsterAfter", func(t *testing.T) {
		sl.InsertAfter(2, 5)
		if outPut() != "5 1 5 2 3 4 5 6 7 8 9 " {
			t.Fatal("fail", outPut())
		}
	})

	t.Run("Remove", func(t *testing.T) {
		sl.Remove(1)
		sl.Remove(2)
		if outPut() != "1 2 3 4 5 6 7 8 9 " {
			t.Fatal("fail", outPut())
		}
	})

	t.Run("Update", func(t *testing.T) {
		sl.Update(2, 99)
		if outPut() != "1 99 3 4 5 6 7 8 9 " {
			t.Fatal("fail", outPut())
		}
	})
}
