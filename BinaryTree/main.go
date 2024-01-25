package main

func main() {
	// Original tree from Wikipedia articl
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
}
