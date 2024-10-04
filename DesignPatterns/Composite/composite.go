package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
)

type Composite struct {
	Name  string
	Elems []Component
}

// Do implements Component.
func (c Composite) Do() string {
	fmt.Printf("%s started jobs\n", c.Name)

	buf := bytes.NewBuffer(nil)
	for _, v := range c.Elems {
		_, err := buf.Write([]byte(v.Do()))
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Printf("%s finished jobs\n", c.Name)
	str, err := buf.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatal(err.Error())
	}

	return str
}
