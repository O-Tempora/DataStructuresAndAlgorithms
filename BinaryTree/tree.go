package main

type Node[T comparable] struct {
	Data  T
	Left  *Node[T]
	Right *Node[T]
}

type BinaryTree[T comparable] struct {
	Head *Node[T]
}
