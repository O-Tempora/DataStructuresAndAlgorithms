package main

import (
	"bufio"
	"fmt"
)

type Adaptee struct{}

// WriteUnadapted needs to be adapted.
func (a Adaptee) WriteUnadapted(buf *bufio.Writer, v []byte) error {
	_, err := buf.Write(v)
	if err != nil {
		return fmt.Errorf("buf.Write: %w", err)
	}

	buf.Flush()

	return nil
}
