package main

import (
	"fmt"
	"time"
)

type Saver interface {
	SaveData(data Data) error
}

type Data struct {
	Name       string    `json:"name"`
	ID         string    `json:"id"`
	Time       time.Time `json:"time"`
	MoneyCents int64     `json:"moneyCents"`
}

type JSONSaver struct {
	validator Validator
	parser    Parser
	writer    FileWriter
}

func NewJSONSaver(p string) (*JSONSaver, error) {
	s := &JSONSaver{
		validator: DataValidator{},
		parser:    JSONParser{},
	}

	writer, err := NewJSONWriter(p)
	if err != nil {
		return nil, fmt.Errorf("failed to create writer: %w", err)
	}
	s.writer = writer

	return s, nil
}

func (s JSONSaver) SaveData(data Data) error {
	if err := s.validator.Validate(data); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	bytes, err := s.parser.ToJSON(data)
	if err != nil {
		return fmt.Errorf("parser error: %w", err)
	}

	if err := s.writer.WriteJSON(bytes); err != nil {
		return fmt.Errorf("write file error: %w", err)
	}

	return nil
}
