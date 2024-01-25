package main

import "fmt"

func main() {
	strStack := Stack[string]{}
	strStack.Initialize(5)

	strStack.Push("Hello")
	strStack.Push("_")
	strStack.Push("World")
	strStack.Push("!")
	strStack.PrintStack()

	fmt.Println("Test peek:", *strStack.Peek())
	fmt.Println("Test pop 1:", *strStack.Pop())
	fmt.Println("Test pop 2:", *strStack.Pop())
	strStack.PrintStack()

	fmt.Println("Empty: ", strStack.IsEmpty(), ", full: ", strStack.IsFull())
}
