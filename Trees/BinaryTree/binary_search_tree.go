package main

import "golang.org/x/exp/constraints"

func binarySearch[T constraints.Ordered](head *Node[T], value T) (res *Node[T]) {
	if head == nil {
		return nil
	}
	switch {
	case head.Data == value:
		return head
	case value < head.Data:
		return binarySearch(head.Left, value)
	case value > head.Data:
		return binarySearch(head.Right, value)
	}
	return nil
}

func binaryInsert[T constraints.Ordered](head *Node[T], value T) *Node[T] {
	if head == nil {
		return &Node[T]{
			Left:  nil,
			Right: nil,
			Data:  value,
		}
	}
	if value < head.Data {
		head.Left = binaryInsert(head.Left, value)
	} else if value > head.Data {
		head.Right = binaryInsert(head.Right, value)
	}
	return head
}

func binaryDelete[T constraints.Ordered](head *Node[T], value T) *Node[T] {
	if head == nil {
		return nil
	}
	if value < head.Data {
		head.Left = binaryDelete(head.Left, value)
	} else if value > head.Data {
		head.Right = binaryDelete(head.Right, value)
	} else {
		var tmp *Node[T]
		switch {
		case head.Left == nil:
			tmp = head.Right
			head = nil
			return tmp
		case head.Right == nil:
			tmp = head.Left
			head = nil
			return tmp
		default:
			tmp = head.Right
			for tmp != nil && tmp.Left != nil {
				tmp = tmp.Left
			}
			head.Data = value
			head.Right = binaryDelete(head.Right, tmp.Data)
		}
	}
	return head
}
