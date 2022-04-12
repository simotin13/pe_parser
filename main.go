package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	filePath := os.Args[1]
	fi, err := os.Stat(filePath)
	if err != nil {
		log.Printf("%s cannot open", filePath)
		os.Exit(-1)
	}

	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("%s cannot open", filePath)
		os.Exit(-1)
	}
	defer f.Close()
	bin := make([]byte, fi.Size())
	f.Read(bin)
}
