package main

import "fmt"

type QuickSort struct{}

// Worst:   N^2
// Average: N*logN
// Best:    N*logN
func (QuickSort) Sort(a []int) []int {
	s := make([]int, len(a))
	copy(s, a)
	length := len(s)
	qsort(s, 0, length-1)
	return s
}

func qsort(a []int, low, high int) {
	if low < high {
		pivot := partition(a, low, high)
		qsort(a, low, pivot-1)
		qsort(a, pivot+1, high)
	}
}

func partition(a []int, low, high int) int {
	pivot := a[high] //Choosing pivot element (first/last/random)
	moving_pointer := low - 1
	for i := low; i < high; i++ {
		if a[i] <= pivot {
			moving_pointer++
			a[i], a[moving_pointer] = a[moving_pointer], a[i]
		}
	}
	a[moving_pointer+1], a[high] = a[high], a[moving_pointer+1]
	fmt.Println(a)
	return moving_pointer + 1
}
