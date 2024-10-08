package main

import (
	"errors"
	"fmt"
	"time"
)

type Validator interface {
	Validate(data Data) error
}

type DataValidator struct{}

func (v DataValidator) Validate(data Data) (err error) {
	if data.Name == "" {
		err = errors.Join(err, fmt.Errorf("empty name"))
	}
	if data.ID == "" {
		err = errors.Join(err, fmt.Errorf("empty id"))
	}
	if data.Time.Equal(time.Time{}) {
		err = errors.Join(err, fmt.Errorf("empty time"))
	}
	if data.MoneyCents < 0 {
		err = errors.Join(err, fmt.Errorf("invalid amount = [%v]", data.MoneyCents))
	}
	return
}
