package main

import "fmt"

type InsertionSort struct{}

// Worst:   N^2 comparisons, N^2 swaps
// Average: N^2 comparisons, N^2 swaps
// Best:    N comparisons, 1 swaps (already sorted)
func (InsertionSort) Sort(a []int) []int {
	s := make([]int, len(a))
	copy(s, a)
	length := len(a)
	pivot_i, pivot_v := 0, 0
	for i := 1; i < length; i++ {
		pivot_i = i - 1
		pivot_v = s[i]
		for pivot_i >= 0 && s[pivot_i] > pivot_v {
			s[pivot_i+1] = s[pivot_i]
			pivot_i--
		}
		s[pivot_i+1] = pivot_v
		fmt.Printf("Step %d: %v\n", i, s)
	}
	return s
}
