package main

import "fmt"

type MergeSort struct{}

// All complexity cases: N*logN
func (MergeSort) Sort(a []int) []int {
	s := make([]int, len(a))
	copy(s, a)
	res := mSort(s)
	fmt.Println(res)
	return s
}

func mSort(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}
	mid := length / 2
	return merge(mSort(a[:mid]), mSort(a[mid:]))
}

func merge(left, right []int) []int {
	fmt.Printf("l:%v r:%v --- ", left, right)
	ll, lr := len(left), len(right) //length left, length right
	i, j, size := 0, 0, ll+lr
	res := make([]int, size)
	for k := 0; k < size; k++ {
		if i > ll-1 && j <= lr-1 { //left is done, right is not
			res[k] = right[j]
			j++
		} else if j > lr-1 && i <= ll-1 { //right is done, left is not
			res[k] = left[i]
			i++
		} else if left[i] < right[j] { //left < right
			res[k] = left[i]
			i++
		} else if left[i] >= right[j] { //right < left
			res[k] = right[j]
			j++
		}
	}
	fmt.Println(res)
	return res
}
