package main

import (
	"errors"
	"fmt"
)

type Queue struct {
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
	return nil
}

func (q *Queue) DeQueue() (interface{}, error) {
	if q.rear == q.front {
		return nil, errors.New("queue is empty")
	}
	item := q.data[q.front]
	q.front, q.data[q.front] = (q.front+1)%q.size, nil
	return item, nil
}

func (q *Queue) OutPut() {
	fmt.Println(q.data)
}

/*
* 维护头部和尾部的下标，
* 入队 尾部循环+1 尾部 = 尾部+1 % 容量   加1 % 容量 == 头部或者尾部的时候，意味着要不队列空了或者满了
* 出队 头部循环+1 头部 = 头部+1 % 容量
 */

func main() {
	q := NewQueue(10)
	for i := 1; i <= 10; i++ {
		q.EnQueue(i)
	}
	q.OutPut()
}
