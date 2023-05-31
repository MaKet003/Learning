package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Printf("minimum one file is needed")
		os.Exit(1)
	}
	for _, v := range os.Args[1:] {
		fd, err := os.Open(v)
		if err != nil {
			log.Println("error open file: ", v)
			os.Exit(2)
		}
		_, err = io.Copy(os.Stdout, fd)
		if err != nil {
			log.Println("error with io.Copy() ", v)
		}
		fd.Close()
	}
}
