package main

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	Data  T
	Left  *Node[T]
	Right *Node[T]
}

type BinaryTree[T constraints.Ordered] struct {
	Head *Node[T]
}
