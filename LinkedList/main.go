package main

func main() {
	n3 := Node[float32]{next: nil, data: 3.14}
	n2 := Node[float32]{next: &n3, data: 2.97}
	n1 := Node[float32]{next: &n2, data: 19.54}
	head := Node[float32]{next: &n1, data: 0.17}
	list := LinkedList[float32]{head: &head}
	list.Traverse()

	n4 := Node[float32]{data: 69.69}
	list.InsertAfter(&n2, &n4)
	list.Traverse()
	n5 := Node[float32]{data: 14.88}
	list.InsertBeginning(&n5)
	list.Traverse()

	list.RemoveBeginning()
	list.RemoveBeginning()
	list.Traverse()
	list.RemoveAfter(&n2)
	list.Traverse()
	list.Search(2.97)
	list.Search(12.34)

	list.Traverse()
	list.Inverse()
	list.Traverse()
}
