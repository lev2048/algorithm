package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	size int
	data []interface{}
}

func NewStack(len int) *Stack {
	return &Stack{
		size: len,
		data: make([]interface{}, 0, len),
	}
}

func (s *Stack) Push(v interface{}) error {
	if len(s.data) == s.size {
		return errors.New("stack is full")
	}
	s.data = append(s.data, v)
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack is empty")
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item, nil
}

func (s *Stack) Print() {
	fmt.Println(s.data, len(s.data), cap(s.data))
}
