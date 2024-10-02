package main

import "log"

func main() {
	var (
		w        = Writer{}
		content  = []byte("write this into file\n")
		filepath = "./res.out"
	)
	if err := write(w, content, filepath); err != nil {
		log.Fatal(err)
	}
}

func write(a Adapter, v []byte, p string) error {
	return a.WriteAdapted(v, p)
}
