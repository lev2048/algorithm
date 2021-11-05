package main

import (
	"fmt"
	"testing"
)

func TestSingleList(t *testing.T) {
	l := new(List)
	data := []int{1, 2, 3, 4, 5, 6, 7}

	outPut := func() string {
		res := ""
		for i := 1; i <= l.Len(); i++ {
			res += fmt.Sprintf("%d ", l.Get(i).Data)
		}
		return res
	}

	t.Run("PushBack", func(t *testing.T) {
		for _, v := range data {
			l.PushBack(v)
		}
		if outPut() != "1 2 3 4 5 6 7 " {
			t.Fatal("fail", outPut())
		}
	})

	t.Run("PushFront", func(t *testing.T) {
		l.PushFront(5)
		if outPut() != "5 1 2 3 4 5 6 7 " {
			t.Fatal("fail", outPut())
		}
	})

	t.Run("InsterAfter", func(t *testing.T) {
		l.Insert(2, 5)
		if outPut() != "5 1 5 2 3 4 5 6 7 " {
			t.Fatal("fail", outPut())
		}
	})

	t.Run("Remove", func(t *testing.T) {
		l.Remove(1)
		l.Remove(2)
		if outPut() != "1 2 3 4 5 6 7 " {
			t.Fatal("fail", outPut())
		}
	})

	t.Run("Update", func(t *testing.T) {
		l.Update(2, 99)
		if outPut() != "1 99 3 4 5 6 7 " {
			t.Fatal("fail", outPut())
		}
	})
}
