package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func (t *BinaryTree[T]) DepthFirstSearch(searchFunc func(*Node[T])) {
	if t.Head == nil {
		return
	}
	searchFunc(t.Head)
	fmt.Println()
}

func traversePreorder[T constraints.Ordered](node *Node[T]) {
	if node == nil {
		return
	}
	fmt.Printf("%c\t", node.Data)
	traversePreorder(node.Left)
	traversePreorder(node.Right)
}

func traverseInorder[T constraints.Ordered](node *Node[T]) {
	if node == nil {
		return
	}
	traverseInorder(node.Left)
	fmt.Printf("%c\t", node.Data)
	traverseInorder(node.Right)
}

func traversePostorder[T constraints.Ordered](node *Node[T]) {
	if node == nil {
		return
	}
	traversePostorder(node.Left)
	traversePostorder(node.Right)
	fmt.Printf("%c\t", node.Data)
}
