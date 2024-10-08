package main

import "encoding/json"

type Parser interface {
	ToJSON(data any) ([]byte, error)
}

type JSONParser struct{}

func (p JSONParser) ToJSON(data any) ([]byte, error) {
	return json.Marshal(data)
}
