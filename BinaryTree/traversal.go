package main

import "fmt"

func (t *BinaryTree[T]) DepthFirstSearch(searchFunc func(*Node[T])) {
	if t.Head == nil {
		return
	}
	searchFunc(t.Head)
	fmt.Println()
}

func traversePreorder[T comparable](node *Node[T]) {
	if node == nil {
		return
	}
	fmt.Printf("%c\t", node.Data)
	traversePreorder(node.Left)
	traversePreorder(node.Right)
}

func traverseInorder[T comparable](node *Node[T]) {
	if node == nil {
		return
	}
	traverseInorder(node.Left)
	fmt.Printf("%c\t", node.Data)
	traverseInorder(node.Right)
}

func traversePostorder[T comparable](node *Node[T]) {
	if node == nil {
		return
	}
	traversePostorder(node.Left)
	traversePostorder(node.Right)
	fmt.Printf("%c\t", node.Data)
}
