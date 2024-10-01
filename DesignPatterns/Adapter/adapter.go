package main

import (
	"bufio"
	"fmt"
	"os"
)

// Adapter interface
type Adapter interface {
	WriteAdapted(v []byte, path string) error
}

// Writer implements Adapter.
type Writer struct{}

func (w Writer) WriteAdapted(v []byte, path string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %w", err)
	}
	defer f.Close()

	buf := bufio.NewWriter(f)
	var a Adaptee
	if err := a.WriteUnadapted(buf, v); err != nil {
		return fmt.Errorf("a.WriteUnadapted: %w", err)
	}

	return nil
}
