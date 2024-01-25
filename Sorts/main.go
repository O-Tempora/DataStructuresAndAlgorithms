package main

import (
	"fmt"
	"log"
	"os"
)

type Sorter interface {
	Sort([]int) []int
}

type Sorting struct {
	sorter Sorter
}

func (s Sorting) Sort(a []int) []int {
	return s.sorter.Sort(a)
}

var sorts = map[string]Sorter{
	"bubble":    BubbleSort{},
	"insertion": InsertionSort{},
	"selection": SelectionSort{},
	"quick":     QuickSort{},
	"merge":     MergeSort{},
}

func main() {
	v, ok := sorts[os.Args[1]]
	if !ok {
		log.Fatal("No such sorting algorithm")
	}

	data := []int{2, 12, 4, 9, 10, -5, 0, 3, -2}
	fmt.Println("Original slice:", data)

	sorting := Sorting{v}
	sorting.Sort(data)
}
