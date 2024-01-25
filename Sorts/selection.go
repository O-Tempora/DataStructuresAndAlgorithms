package main

import "fmt"

type SelectionSort struct{}

// Worst:   N^2 comparisons, N swaps
// Average: N^2 comparisons, N swaps
// Best:    N^2 comparisons, 1 swaps (already sorted)
func (SelectionSort) Sort(a []int) []int {
	s := make([]int, len(a))
	copy(s, a)
	length := len(s)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		if min != i {
			s[min], s[i] = s[i], s[min]
		}
		fmt.Printf("Step %d: %v\n", i, s)
	}

	return s
}
