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
