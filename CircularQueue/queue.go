package main

import "fmt"

type Queue[T comparable] struct {
	head int
	tail int
	size int
	arr  []T
}

func (q *Queue[T]) Enqueue(element T) {
	if !q.IsFull() {
		if q.head == -1 {
			q.head = 0
		}
		q.tail = (q.tail + 1) % q.size
		q.arr[q.tail] = element
	}
}
func (q *Queue[T]) Dequeue() *T {
	if q.IsEmpty() {
		return nil
	}
	top := q.arr[q.head]
	if q.head == q.tail {
		q.head, q.tail = -1, -1
	} else {
		q.head = (q.head + 1) % q.size
	}
	return &top
}
func (q *Queue[T]) Initialize(size int) {
	q.size = size
	q.head = -1
	q.tail = -1
	q.arr = make([]T, size)
}
func (q Queue[T]) IsEmpty() bool {
	return q.head == -1
}
func (q Queue[T]) IsFull() bool {
	return (q.tail+1)%q.size == q.head
}
func (q Queue[T]) PrintQueue() {
	fmt.Println("Head: ", q.head, ", tail: ", q.tail)
	fmt.Println(q.arr)
}
