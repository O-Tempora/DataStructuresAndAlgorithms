package main

import "fmt"

type Stack[T comparable] struct {
	size int
	arr  []T
	top  int
}

func (st *Stack[T]) Push(element T) {
	if !st.IsFull() {
		st.top++
		st.arr[st.top] = element
	}
}
func (st *Stack[T]) Pop() *T {
	if !st.IsEmpty() {
		top := st.arr[st.top]
		st.top -= 1
		return &top
	}
	return nil
}
func (st Stack[T]) Peek() *T {
	if !st.IsEmpty() {
		return &st.arr[st.top]
	}
	return nil
}
func (st Stack[T]) IsEmpty() bool {
	return st.top < 0
}
func (st Stack[T]) IsFull() bool {
	return st.top == st.size-1
}
func (st *Stack[T]) Initialize(size int) {
	st.size = size
	st.top = -1
	st.arr = make([]T, size)
}
func (st Stack[T]) PrintStack() {
	if !st.IsEmpty() {
		for i := 0; i <= st.top; i++ {
			fmt.Print(st.arr[i])
		}
		fmt.Println()
	}
}
