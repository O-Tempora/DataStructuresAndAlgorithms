package main

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}
type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Traverse() {
	node := l.head
	for node != nil {
		fmt.Printf("%v ", node.data)
		node = node.next
	}
	fmt.Println()
}
func (LinkedList[T]) InsertAfter(node, newNode *Node[T]) {
	newNode.next = node.next
	node.next = newNode
}
func (l *LinkedList[T]) InsertBeginning(newNode *Node[T]) {
	newNode.next = l.head
	l.head = newNode
}
func (LinkedList[T]) RemoveAfter(node *Node[T]) {
	obsolete := node.next
	node.next = node.next.next
	obsolete.next = nil
}
func (l *LinkedList[T]) RemoveBeginning() {
	obsolete := l.head
	l.head = l.head.next
	obsolete.next = nil
}
func (l *LinkedList[T]) Search(element T) {
	node := l.head
	for node != nil {
		if node.data == element {
			fmt.Println("Element", element, "exists")
			return
		}
		node = node.next
	}
	fmt.Println("Element", element, "does not exist")
}
func (l *LinkedList[T]) Inverse() {
	curr := l.head
	var prev, next *Node[T] = nil, nil
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}
