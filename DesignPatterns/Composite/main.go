package main

import "fmt"

func main() {
	cmps := Composite{
		Name: "MainComposite",
		Elems: []Component{
			Leaf{
				Name: "MainLeaf1",
			},
			Composite{
				Name: "SecondaryComposite",
				Elems: []Component{
					Leaf{
						Name: "SecondaryLeaf1",
					},
					Leaf{
						Name: "SecondaryLeaf2",
					},
				},
			},
		},
	}

	PrintResult(cmps)
}

func PrintResult(c Component) {
	fmt.Printf("Resulting string: %s\n", c.Do())
}
