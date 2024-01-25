package main

import "fmt"

type BubbleSort struct{}

// Worst:   N^2 comparisons, N^2 swaps
// Average: N^2 comparisons, N^2 swaps
// Best:    N comparisons, 1 swaps (already sorted)
func (BubbleSort) Sort(a []int) []int {
	s := make([]int, len(a))
	copy(s, a)
	length := len(s)
	var swapped bool
	for i := 0; i < length-1; i++ {
		swapped = false
		for j := 0; j < length-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		fmt.Printf("Step %d: %v\n", i, s)
	}
	return s
}
