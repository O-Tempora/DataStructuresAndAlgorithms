package main

import (
	"bufio"
	"fmt"
	"os"
)

type FileWriter interface {
	WriteJSON(data []byte) error
	SetData(b []byte)
}

type JSONWriter struct {
	data   []byte
	buffer *bufio.Writer
	path   string
}

func (w *JSONWriter) SetData(b []byte) {
	w.data = b
}

func NewJSONWriter(p string) (*JSONWriter, error) {
	wr := &JSONWriter{
		path: p,
	}

	f, err := os.OpenFile(p, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", p, err)
	}
	wr.buffer = bufio.NewWriter(f)

	return wr, nil
}

func (w JSONWriter) WriteJSON(data []byte) error {
	data = append(data, '\n')
	w.SetData(data)
	_, err := w.buffer.Write(w.data)
	if err != nil {
		return fmt.Errorf("failed to write to buffer: %w", err)
	}

	return w.buffer.Flush()
}
