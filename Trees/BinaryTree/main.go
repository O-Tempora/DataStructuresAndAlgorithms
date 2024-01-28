package main

import "fmt"

func main() {
	// Original tree from Wikipedia article
	tree := BinaryTree[rune]{
		Head: &Node[rune]{
			Data: 'F',
			Left: &Node[rune]{
				Data: 'B',
				Left: &Node[rune]{
					Data:  'A',
					Left:  nil,
					Right: nil,
				},
				Right: &Node[rune]{
					Data: 'D',
					Left: &Node[rune]{
						Data:  'C',
						Left:  nil,
						Right: nil,
					},
					Right: &Node[rune]{
						Data:  'E',
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: &Node[rune]{
				Data: 'G',
				Left: nil,
				Right: &Node[rune]{
					Data: 'I',
					Left: &Node[rune]{
						Data:  'H',
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}

	tree.DepthFirstSearch(traversePreorder)
	tree.DepthFirstSearch(traverseInorder)
	tree.DepthFirstSearch(traversePostorder)

	fmt.Println(binarySearch[rune](tree.Head, 'E'))
	fmt.Println(binarySearch[rune](tree.Head, 'S'))
	fmt.Println(binarySearch[rune](tree.Head, 'F'))
	fmt.Println(binarySearch[rune](tree.Head, 'X'))

	binaryInsert[rune](tree.Head, 'K')
	binaryInsert[rune](tree.Head, 'J')
	tree.DepthFirstSearch(traverseInorder)

	fmt.Println(binaryDelete[rune](tree.Head, 'S'))
	fmt.Println(binaryDelete[rune](tree.Head, 'A'))
	tree.DepthFirstSearch(traverseInorder)
}
