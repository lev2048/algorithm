package main

import (
	"testing"
)

var queue *Queue

func init() {
	queue = NewQueue(100)
}

func Benchmark_Queue(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		queue.EnQueue(i)
		queue.DeQueue()
	}
}
