package queue

import (
	"errors"
	"fmt"
)

/*
* 维护头部和尾部的下标，
* 入队 尾部循环+1 尾部 = 尾部+1 % 容量   加1 % 容量 == 头部或者尾部的时候，意味着要不队列空了或者满了
* 出队 头部循环+1 头部 = 头部+1 % 容量
 */

type Queue struct {
	len   int
	size  int
	rear  int
	front int
	data  []interface{}
}

func NewQueue(size int) *Queue {
	return &Queue{
		len:   0,
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
	q.data[q.rear], q.rear = v, (q.rear+1)%q.size
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

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) OutPut() {
	fmt.Println(q.data)
}
