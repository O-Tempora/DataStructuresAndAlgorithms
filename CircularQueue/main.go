package main

import "fmt"

func main() {
	qu := Queue[rune]{}
	qu.Initialize(4)

	qu.Enqueue('h')
	qu.Enqueue('e')
	qu.Enqueue('l')
	qu.Enqueue('p')
	qu.PrintQueue()

	fmt.Println("Test Dequeue 1:", *qu.Dequeue())
	qu.PrintQueue()
	fmt.Println("Test Dequeue 1:", *qu.Dequeue())
	qu.PrintQueue()

	qu.Enqueue('p')
	qu.PrintQueue()
}
