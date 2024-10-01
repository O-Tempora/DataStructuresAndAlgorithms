package main

import "log"

func main() {
	var (
		w        = Writer{}
		content  = []byte("write this into file\n")
		filepath = "./out"
	)
	if err := w.WriteAdapted(content, filepath); err != nil {
		log.Fatal(err)
	}
}
