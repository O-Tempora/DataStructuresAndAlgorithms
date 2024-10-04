package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Leaf struct {
	Name string
}

// Do implements Component.
func (l Leaf) Do() string {
	str := uuid.NewString()
	fmt.Printf("%s leaf generated str\n", l.Name)
	return str
}
