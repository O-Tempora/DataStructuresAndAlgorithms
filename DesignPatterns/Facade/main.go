package main

import (
	"log"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
)

const (
	path = "save.log"
)

func main() {
	saver, err := NewJSONSaver(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	data := Data{
		Name:       "Some_Name",
		ID:         uuid.NewString(),
		Time:       time.Now(),
		MoneyCents: rand.Int64N(10_000_000),
	}

	SaveDataToFile(saver, data)
}

func SaveDataToFile(saver Saver, data Data) error {
	return saver.SaveData(data)
}
